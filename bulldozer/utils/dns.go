package utils

import "github.com/dollarkillerx/easy_dns"

func SearchDns(domain string, dns string) ([]string, error) {
	lookupIP, err := easy_dns.LookupIP(domain, dns)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, v := range lookupIP.Answers {
		if v.Header.Type == easy_dns.TypeA {
			result = append(result, v.Body.GoString())
		}
	}

	return result, nil
}
