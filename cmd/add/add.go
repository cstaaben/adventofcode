package add

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/cstaaben/adventofcode/internal/convert"
)

const (
	addYearTemplateFile   = `add_year.gotmpl`
	parentCmdTemplateFile = `parent_cmd.gotmpl`
	dayCmdTemplateFile    = `day_cmd.gotmpl`
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add_year",
		Short: "add a new package for another year of the Advent of Code",
		RunE:  addYear,
		Args:  cobra.ExactArgs(1),
	}
}

func addYear(_ *cobra.Command, args []string) error {
	templateArgs := make(map[string]interface{})

	t, err := template.ParseGlob("templates/*.gotmpl")
	if err != nil {
		return fmt.Errorf("parsing template files: %w", err)
	}

	var year int
	year, err = strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("parsing argument: %w", err)
	}

	yearWord := convert.NumberToWord(year)
	yearPkg := strings.ToLower(strings.ReplaceAll(yearWord, "-", ""))
	templateArgs["LowerYear"] = yearPkg

	if err = os.Mkdir(fmt.Sprintf("./internal/%s", yearPkg), 0755); err != nil {
		return fmt.Errorf("creating new directory: %w", err)
	}

	err = executeTemplateToFile(t.Lookup(addYearTemplateFile), fmt.Sprintf("./cmd/%s.go", yearPkg), templateArgs)
	if err != nil {
		return fmt.Errorf("add year template: %w", err)
	}

	err = executeTemplateToFile(t.Lookup(parentCmdTemplateFile), fmt.Sprintf("./internal/%s/%s.go", yearPkg, yearPkg), templateArgs)
	if err != nil {
		return fmt.Errorf("parent command template: %w", err)
	}

	templateArgs["YearInt"] = year
	for i := 0; i < 25; i++ {
		day := i + 1
		dayWord := convert.NumberToWord(day)
		dayPkg := fmt.Sprintf("day%d", day)

		templateArgs["Package"] = dayPkg
		templateArgs["LowerDay"] = strings.ToLower(strings.ReplaceAll(dayWord, "-", ""))
		templateArgs["CapitalDay"] = strings.Title(dayWord)
		templateArgs["PascalCaseDay"] = strings.ReplaceAll(strings.Title(strings.Join(strings.Split(dayWord, "-"), " ")), " ", "")
		templateArgs["DayInt"] = day

		err = os.Mkdir(fmt.Sprintf("./internal/%s/%s", yearPkg, dayPkg), 0755)
		if err != nil {
			return fmt.Errorf("making command directory for day %d: %w", day, err)
		}

		err = executeTemplateToFile(t.Lookup(dayCmdTemplateFile), fmt.Sprintf("./internal/%s/%s/cmd.go", yearPkg, dayPkg), templateArgs)
		if err != nil {
			return fmt.Errorf("creating day %d command: %w", day, err)
		}
	}

	return nil
}

func executeTemplateToFile(t *template.Template, filename string, args map[string]interface{}) error {
	if t == nil {
		return errors.New("unable to find template")
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("opening file for template: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	err = t.Execute(writer, args)
	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}
	if err = writer.Flush(); err != nil {
		return fmt.Errorf("flushing file contents: %w", err)
	}

	return nil
}
