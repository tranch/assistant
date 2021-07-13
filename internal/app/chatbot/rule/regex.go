package rule

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

type Rule struct {
	Regex string
	Help  string
	Parse func(context.Context, rulebot.IComponent, string, []string) []string
}

type regexRuleset struct {
	regexes map[string]*template.Template
	rules   []Rule
}

func (r regexRuleset) Name() string {
	return "Regex Ruleset"
}

func (r regexRuleset) Boot(_ *rulebot.RuleBot) {}

func (r regexRuleset) HelpRule(b *rulebot.RuleBot, _ string) string {
	botName := b.Name()
	var helpMsg string
	for _, rule := range r.rules {
		var finalRegex bytes.Buffer
		_ = r.regexes[rule.Regex].Execute(&finalRegex, struct{ RobotName string }{botName})

		helpMsg = fmt.Sprintln(helpMsg, finalRegex.String(), "-", rule.Help)
	}
	return strings.TrimLeftFunc(helpMsg, unicode.IsSpace)
}

func (r regexRuleset) ParseRule(ctx context.Context, b *rulebot.RuleBot, in string) []string {
	for _, rule := range r.rules {
		botName := b.Name()
		var finalRegex bytes.Buffer
		if _, ok := r.regexes[rule.Regex]; !ok {
			r.regexes[rule.Regex] = template.Must(template.New(rule.Regex).Parse(rule.Regex))
		}
		_ = r.regexes[rule.Regex].Execute(&finalRegex, struct{ RobotName string }{botName})
		sanitizedRegex := strings.TrimSpace(finalRegex.String())
		re := regexp.MustCompile(sanitizedRegex)
		matched := re.MatchString(in)
		if !matched {
			continue
		}

		args := re.FindStringSubmatch(in)
		if ret := rule.Parse(ctx, b.Comp, in, args); len(ret) > 0 {
			return ret
		}
	}

	return []string{}
}

func New(rules []Rule) *regexRuleset {
	r := &regexRuleset{
		regexes: make(map[string]*template.Template),
		rules:   rules,
	}
	for _, rule := range rules {
		r.regexes[rule.Regex] = template.Must(template.New(rule.Regex).Parse(rule.Regex))
	}
	return r
}
