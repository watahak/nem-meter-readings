package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"nem12/models"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ParseRecords handles the file according to specs as defined: https://aemo.com.au/-/media/files/electricity/nem/retail_and_metering/market_settlement_and_transfer_solutions/2022/mdff-specification-nem12-nem13-v25.pdf?la=en
//   - interval length -> ninth value in the 200 records, where
//   - nmi ->  2nd value of 200 e.g. NEM1201009
//   - timestamp -> second value in the 300 records
//   - consumption -> sum of values of the interval length i.e index 2 to length - 1
func ParseRecord(reader *csv.Reader) ([]*models.MeterReading, error) {
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	var currentNMI *string
	var currentIntervalLength *int
	recordLine := 0
	var meterReadings []*models.MeterReading

	// Read and parse the CSV file line by line
	for {
		record, err := reader.Read()
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			log.Printf(`failed to read file`)
			return nil, fmt.Errorf(`failed to read file`)
		}

		recordLine++

		switch recordIndicator := record[0]; recordIndicator {
		case "100":
		case "200":
			currentNMI = &record[1]
			intervalLength, err := strconv.Atoi(record[8])
			if err != nil {
				log.Printf(`failed to convert interval length to integer at line:%d`, recordLine)
				return nil, fmt.Errorf(`failed to convert interval length to integer at line:%d`, recordLine)
			}
			currentIntervalLength = &intervalLength

		case "300":
			if currentIntervalLength == nil || currentNMI == nil {
				log.Printf(`failed locate interval length and nmi values in 200 records in record:%d`, recordLine)
				return nil, fmt.Errorf(`failed locate interval length and nmi values in 200 records in record:%d`, recordLine)
			}

			var sum float64

			firstIndex := 2
			last := 2 + (24*60) / *currentIntervalLength

			date, err := time.Parse("20060102", record[1])
			if err != nil {
				log.Printf(`failed to convert timestamp at line:%d`, recordLine)
				return nil, fmt.Errorf(`failed to convert timestamp at line:%d`, recordLine)
			}

			for i := firstIndex; i < last; i++ {
				float64Val, err := strconv.ParseFloat(strings.TrimSpace(record[i]), 32)
				if err != nil {
					log.Printf(`failed to convert consumption at line:%d, index:%d`, recordLine, i)
					return nil, fmt.Errorf(`failed to convert consumption at line:%d, index:%d`, recordLine, i)
				}
				sum += float64Val
			}

			meterReading := &models.MeterReading{
				ID:          uuid.New(),
				NMI:         *currentNMI,
				Timestamp:   date,
				Consumption: float32(int(sum*1000)) / 1000,
			}

			meterReadings = append(meterReadings, meterReading)

		case "400":
		case "500":
		case "900":
		default:
			log.Print("unknown record: ", record)
		}
	}

	return meterReadings, nil
}
