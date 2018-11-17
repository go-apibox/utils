package utils

import (
	"path/filepath"
	"strings"
)

type Matcher struct {
	whiteList     []string
	whiteMap      map[string]bool
	whitePatterns []string
	blackList     []string
	blackMap      map[string]bool
	blackPatterns []string
}

func NewMatcher() *Matcher {
	m := new(Matcher)
	m.whiteList = make([]string, 0)
	m.whiteMap = make(map[string]bool, 0)
	m.whitePatterns = make([]string, 0)
	m.blackList = make([]string, 0)
	m.blackMap = make(map[string]bool, 0)
	m.blackPatterns = make([]string, 0)
	return m
}

func (m *Matcher) SetWhiteList(whiteList []string) *Matcher {
	m.whiteList = whiteList
	m.whiteMap = make(map[string]bool, 0)
	m.whitePatterns = make([]string, 0)
	for _, item := range whiteList {
		if strings.Index(item, "*") >= 0 {
			m.whitePatterns = append(m.whitePatterns, item)
		} else {
			m.whiteMap[item] = true
		}
	}
	return m
}

func (m *Matcher) AddWhiteItem(item string) *Matcher {
	for _, v := range m.whiteList {
		if item == v {
			return m
		}
	}
	m.whiteList = append(m.whiteList, item)
	return m
}

func (m *Matcher) DelWhiteItem(item string) *Matcher {
	for i, v := range m.whiteList {
		if item == v {
			m.whiteList = append(m.whiteList[0:i], m.whiteList[i+1:]...)
			return m
		}
	}
	return m
}

func (m *Matcher) SetBlackList(blackList []string) *Matcher {
	m.blackList = blackList
	m.blackMap = make(map[string]bool, 0)
	m.blackPatterns = make([]string, 0)
	for _, item := range blackList {
		if strings.Index(item, "*") >= 0 {
			m.blackPatterns = append(m.blackPatterns, item)
		} else {
			m.blackMap[item] = true
		}
	}
	return m
}

func (m *Matcher) AddBlackItem(item string) *Matcher {
	for _, v := range m.blackList {
		if item == v {
			return m
		}
	}
	m.blackList = append(m.blackList, item)
	return m
}

func (m *Matcher) DelBlackItem(item string) *Matcher {
	for i, v := range m.blackList {
		if item == v {
			m.blackList = append(m.blackList[0:i], m.blackList[i+1:]...)
			return m
		}
	}
	return m
}

func (m *Matcher) Match(target string) bool {
	// check whitelist
	matchWhite := false
	if len(m.whiteMap) == 0 && len(m.whitePatterns) == 0 {
		return false
	} else {
		_, exists := m.whiteMap[target]
		if exists {
			matchWhite = true
		} else {
			// check patterns
			for _, pattern := range m.whitePatterns {
				matched, _ := filepath.Match(pattern, target)
				if matched {
					matchWhite = true
					break
				}
			}
		}
	}
	if !matchWhite {
		return false
	}

	// check blacklist
	_, exists := m.blackMap[target]
	if exists {
		return false
	}
	// check patterns
	for _, pattern := range m.blackPatterns {
		matched, _ := filepath.Match(pattern, target)
		if matched {
			return false
		}
	}

	return true
}
