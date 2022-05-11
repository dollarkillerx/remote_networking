package utils

import (
	"github.com/dollarkillerx/easy_dns"
	"github.com/patrickmn/go-cache"

	"time"
)

var dnsCache *cache.Cache

func init() {
	dnsCache = cache.New(5*time.Minute, 10*time.Minute)
}

func SearchDns(domain string, dns string) ([]string, error) {

	pDns, found := dnsCache.Get(domain)
	if found {
		pds := pDns.(string)
		if pds != "" {
			return []string{pds}, nil
		}
	}

	ds, err := lookDns(domain, dns)

	if err == nil {
		if len(ds) != 0 {
			dnsCache.Set(domain, ds[0], time.Hour)
		}
	}

	return ds, err
}

func lookDns(domain string, dns string) ([]string, error) {
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
