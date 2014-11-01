package textgameslib

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Option struct {
	Description string
	Flags       []string
	Id          int
}

func (o *Option) String() string {
	return o.Description
}

type Chooser struct {
	listOfOption []Option
	altOption    map[string]Option
}

func (c *Chooser) Add(option Option) bool {
	c.listOfOption = append(c.listOfOption, option)
	for _, flag := range option.Flags {
		c.altOption[flag] = option
	}
	return true
}

func (c *Chooser) Ask() string {
	var result string
	for index, option := range c.listOfOption {
		result += strconv.Itoa(index+1) + " : " + option.String() + "\n"
	}
	return result
}

func (c *Chooser) RandOption() Option {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := random.Intn(len(c.listOfOption))

	return c.listOfOption[index]
}

func (c *Chooser) Answer(input string) (Option, error) {
	var result Option

	if index, err := strconv.Atoi(input); err == nil && index < len(c.listOfOption)+1 {
		index--
		result = c.listOfOption[index]
	} else if option, ok := c.altOption[input]; ok {
		result = option
	} else {
		return result, errors.New("Chooser: " + input + " is not an option.")
	}
	return result, nil
}

func NewChooser() *Chooser {
	return &Chooser{altOption: make(map[string]Option)}
}
