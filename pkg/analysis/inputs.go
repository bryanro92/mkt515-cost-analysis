package analysis

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s *situation) populateQuestions() {
	s.questions = []string{
		"Enter fixed cost: ",
		"Enter variable cost per unit: ",
		"Enter total target market (ttl number of customers / year): ",
		"Enter decimal of market awareness: ",
		"Enter average revenue per customer: ",
	}
}

func (s *situation) dataInit() error {
	var err error
	s.data.fixedCost, err = strconv.Atoi(s.answers[0])
	if err != nil {
		return err
	}
	s.data.variableCost, err = strconv.Atoi(s.answers[1])
	if err != nil {
		return err
	}
	s.data.totalTargetMarket, err = strconv.Atoi(s.answers[2])
	if err != nil {
		return err
	}
	s.data.marketAwareness, err = strconv.ParseFloat(s.answers[3], 64)
	if err != nil {
		return err
	}
	s.data.avgRevenuePerCx, err = strconv.Atoi(s.answers[4])
	return err
}

func (s *situation) init() error {
	s.populateQuestions()

	err := s.parseInput()
	if err != nil {
		return err
	}

	return s.dataInit()
}

func (s *situation) parseInput() error {
	reader := bufio.NewReader(os.Stdin)
	for _, q := range s.questions {
		fmt.Print(q)
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.Trim(input, "\n")
		s.answers = append(s.answers, input)
	}
	return nil
}
