package lib

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func GetUniqueTrxId() string {
	uid := strings.ToUpper(uuid.New().String()[0:6])
	prefix := strconv.Itoa(time.Now().YearDay())
	postfix := strconv.Itoa(time.Now().Year() % 1000)
	uniqueId := prefix + uid + postfix
	fmt.Printf("GetUniqueTrxId:      %s\n", uniqueId)
	return uniqueId
}

func GetUniqueInvoiceNumber(invoicePrefix string) string {
	uid := strings.ToUpper(uuid.New().String()[0:6])
	prefix := strconv.Itoa(time.Now().YearDay())
	postfix := strconv.Itoa(time.Now().Year() % 1000)
	invoiceNumber := invoicePrefix + prefix + uid + postfix
	fmt.Printf("GetUniqueInvoiceNumber:      %s\n", invoiceNumber)
	return invoiceNumber
}

func GenerateFileName(filename string) string {
	t := time.Now()
	_ = t.String()
	filename = strings.TrimSpace(filename)
	filename = strings.ReplaceAll(filename, " ", "")
	filename = strings.ReplaceAll(filename, "-", "_")
	currentDatetime := t.Format("02_01_2006_15_04_05PM")
	uid := strings.ToUpper(uuid.New().String()[0:6])
	newfilename := currentDatetime + "_" + uid + "_" + filename
	return newfilename
}

func GenerateSlug(name string) string {
	slugText := slug.MakeLang(name, "en")
	return slugText
}

func ReplaceS3UrlWithCloudfrontUrl(url, AwsS3BaseURL, AwsCloudfrontURL string) string {
	return strings.Replace(url, AwsS3BaseURL, AwsCloudfrontURL, 1)
}

func GetBaseUrlFromS3(url, AwsS3BaseURL, AwsCloudfrontURL string) (string, error) {
	if strings.Contains(url, AwsCloudfrontURL) {
		url = url[len(AwsCloudfrontURL)+1:]
	} else if strings.Contains(url, AwsS3BaseURL) {
		url = url[len(AwsS3BaseURL)+1:]
	} else {
		return url, fmt.Errorf("custom_error!!!invalid url")
	}
	fmt.Println(url)
	return url, nil
}
