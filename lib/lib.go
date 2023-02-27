package lib

import (
	crand "crypto/rand"
	_ "embed"
	"encoding/base64"
	"html/template"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	Pkg    string
	Func   string
	Key    string
	Input  string
	Output string
	TmpVar string
	Var    string
	Src    string
}

//go:embed template.tmpl
var tmpl string

func KeyGen() string {
	key := make([]byte, 32)
	_, _ = crand.Read(key)
	return base64.RawStdEncoding.EncodeToString(key)
}

func StrGen(length int) string {
	const charset = "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ_0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	_, err := strconv.Atoi(string(b[0]))
	if err != nil {
		return string(b)
	} else {
		return StrGen(length)
	}
}

func IsValidVariableName(s string) bool {
	regex := regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")
	return regex.MatchString(s)
}

func FileNameWithoutExtension(fileName string) string {
	return filepath.Base(strings.TrimSuffix(fileName, filepath.Ext(fileName)))
}

func Run(cfg Config) {
	data, err := os.ReadFile(cfg.Input)
	if err != nil {
		log.Fatal("Missing input file")
	}

	if !IsValidVariableName(cfg.Var) {
		log.Fatal("Invalid variable name")
	}

	out, err := os.Create(cfg.Output)
	if err != nil {
		log.Fatal("Cannot create file")
	}
	key, err := base64.RawStdEncoding.DecodeString(cfg.Key)
	if err != nil {
		log.Fatal("Invalid key")
	}
	_, _ = out.Write(Compress(data, key))
	out.Close()

	srcf, err := os.Create(cfg.Src)
	if err != nil {
		log.Fatal("Cannot create file")
	}
	defer srcf.Close()

	src, err := template.New(cfg.Src).Parse(tmpl)
	if err != nil {
		log.Fatal("Cannot parse text")
	}
	err = src.Execute(srcf, cfg)
	if err != nil {
		log.Fatal("Cannot write file")
	}
}
