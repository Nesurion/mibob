package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/nesurion/go-limitless"
)

var (
	controller *limitless.LimitlessController
)

type Config struct {
	Bridge string `json:"bridge"`
}

func main() {
	flag.Parse()

	for {
		c := parseRGB()
		sendColor(c)
	}
}

func init() {
	conf, err := loadConfig()
	if err != nil {
		panic(err)
	}

	controller, err = limitless.NewLimitlessController(conf.Bridge)
	if err != nil {
		panic(err)
	}

	group := limitless.LimitlessGroup{
		Id:         0,
		Controller: controller,
	}
	controller.Groups = []limitless.LimitlessGroup{group}

	controller.Groups[0].On()
}

func parseRGB() colorful.Color {
	reader := bufio.NewReader(os.Stdin)
	rawRGB, _ := reader.ReadString('\n')
	rawRGB = strings.TrimSuffix(rawRGB, "\n")
	rgbString := strings.Split(rawRGB, " ")

	c := colorful.Color{}

	c.R = toFloat(rgbString[0])
	c.G = toFloat(rgbString[1])
	c.B = toFloat(rgbString[2])

	return c
}

func toFloat(c string) float64 {
	f, _ := strconv.ParseFloat(c, 64)
	return f
}

func sendColor(color colorful.Color) {
	h, _, _ := color.Hsv()
	h = 240.0 - h
	if h < 0 {
		h = 360.0 + h
	}
	scaled_h := uint8(h * 255.0 / 360.0)

	controller.Groups[0].SendColorByte(scaled_h)
}

func loadConfig() (Config, error) {
	c := Config{}
	absConfigPath, err := filepath.Abs("mibob.config")
	if err != nil {
		return c, err
	}
	file, _ := os.Open(absConfigPath)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}
