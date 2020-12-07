package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func validateYear(s string, start, end int) bool {
	y, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if y >= start && y <= end {
		return true
	}
	return false
}

type passport struct {
	ecl, pid, eyr, hcl, byr, iyr, cid, hgt string
}

func (p *passport) parseLine(l string) {
	pairs := strings.Split(l, " ")
	for _, pair := range pairs {
		pairComponents := strings.Split(pair, ":")
		switch pairComponents[0] {
		case "ecl":
			p.ecl = pairComponents[1]
		case "pid":
			p.pid = pairComponents[1]
		case "eyr":
			p.eyr = pairComponents[1]
		case "hcl":
			p.hcl = pairComponents[1]
		case "byr":
			p.byr = pairComponents[1]
		case "iyr":
			p.iyr = pairComponents[1]
		case "cid":
			p.cid = pairComponents[1]
		case "hgt":
			p.hgt = pairComponents[1]
		default:
			log.Fatalf("got unexpected key for passport: %s", pairComponents[0])
		}
	}
}

func (p *passport) validSimple() bool {
	return p.ecl != "" &&
		p.pid != "" &&
		p.eyr != "" &&
		p.hcl != "" &&
		p.byr != "" &&
		p.iyr != "" &&
		p.hgt != ""
}

func (p *passport) validateByr() bool {
	return validateYear(p.byr, 1920, 2002)
}

func (p *passport) validateIyr() bool {
	return validateYear(p.iyr, 2010, 2020)
}

func (p *passport) validateEyr() bool {
	return validateYear(p.eyr, 2020, 2030)
}

func (p *passport) validateHgt() bool {
	letters := []rune(p.hgt)
	if !(len(letters) == 5 || len(letters) == 4) {
		return false
	}

	suffix := string(letters[len(letters)-2:])
	if !(suffix == "cm" || suffix == "in") {
		return false
	}

	height, err := strconv.Atoi(string(letters[:len(letters)-2]))
	if err != nil {
		return false
	}
	if suffix == "cm" && (height < 150 || height > 193) {
		return false
	}
	if suffix == "in" && (height < 59 || height > 76) {
		return false
	}

	return true
}

func (p *passport) validateHcl() bool {
	letters := []rune(p.hcl)
	if len(letters) != 7 {
		return false
	}
	if letters[0] != '#' {
		return false
	}
	for _, l := range letters[1:] {
		ord := int(l)
		if !(ord >= 48 && ord <= 57) && !(ord >= 97 && ord <= 102) {
			return false
		}
	}
	return true
}

func (p *passport) validateEcl() bool {
	switch p.ecl {
	case "amb":
		fallthrough
	case "blu":
		fallthrough
	case "brn":
		fallthrough
	case "gry":
		fallthrough
	case "grn":
		fallthrough
	case "hzl":
		fallthrough
	case "oth":
		return true
	}
	return false
}

func (p *passport) validatePid() bool {
	letters := []rune(p.pid)
	if len(letters) != 9 {
		return false
	}
	_, err := strconv.Atoi(p.pid)
	if err != nil {
		return false
	}
	return true
}

func (p *passport) validComplex() bool {
	return p.validateByr() &&
		p.validateIyr() &&
		p.validateEyr() &&
		p.validateHgt() &&
		p.validateHcl() &&
		p.validateEcl() &&
		p.validatePid()
}

func first(passports []passport) {
	var count int
	for _, p := range passports {
		if p.validSimple() {
			count++
		}
	}
	log.Println(count)
}

func second(passports []passport) {
	var count int
	for _, p := range passports {
		if p.validComplex() {
			count++
		}
	}
	log.Println(count)
}

func main() {
	f, err := os.Open("./4.txt")
	if err != nil {
		log.Fatalf("unable to open file %+v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var passports []passport

ScanLoop:
	for scanner.Scan() {
		p := passport{}
		for {
			s := scanner.Text()
			if len(s) == 0 {
				passports = append(passports, p)
				break
			}
			p.parseLine(s)
			if !scanner.Scan() {
				passports = append(passports, p)
				break ScanLoop
			}
		}

	}

	first(passports)
	second(passports)
}
