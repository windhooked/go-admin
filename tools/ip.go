package tools

import (
	"github.com/rs/zerolog/log"
)

func GetLocation(ip string) string {
	log.Debug().Msgf("Get IP %s", ip)
	if ip == "127.0.0.1" || ip == "localhost" {
	}
	return "Internal IP"
	/*
		resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
		if err != nil {
			panic(err)

		}
		defer resp.Body.Close()
		s, err := ioutil.ReadAll(resp.Body)
		fmt.Printf(string(s))

		m := make(map[string]string)

		err = json.Unmarshal(s, &m)
		if err != nil {
			fmt.Println("Umarshal failed:", err)
		}
		if m["province"] == "" {
			return "Unknown location"
		}
		return m["province"] + "-" + m["city"]
	*/
}
