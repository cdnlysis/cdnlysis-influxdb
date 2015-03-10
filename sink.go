package main

import (
	"log"

	"github.com/cdnlysis/cdnlysis"
	"github.com/cdnlysis/cdnlysis_engine/client"
)

var cachedInfluxConn *client.Client

func makeInfluxSession() {
	conn, err := client.New(&Settings.Influx)
	if err != nil {
		log.Println("Cannot connect to Influx", err)
		return
	}

	cachedInfluxConn = conn
}

func RefreshInfluxSession() {
	log.Println("Refreshing Influx Connection")
	makeInfluxSession()
}

var columns = []string{"date", "x_edge_location", "sc_bytes", "c_ip",
	"cs_method", "cs_host_", "cs_uri_stem", "sc_status", "cs_uri_query",
	"x_edge_result_type", "x_edge_request_id", "x_host_header",
	"cs_protocol", "cs_bytes", "time_taken", "cs_referer_", "cs_user_agent_",
	"cs_cookie_", "time",
}

func makeRecord(record *cdnlysis.LogRecord) *client.Series {
	record.Convert()

	series := client.Series{
		"cdn",
		columns,
		nil,
	}

	new_array := []interface{}{
		record.Date,
		record.EdgeLocation,
		record.BytesSent,
		record.IP,
		record.Method,
		record.Host,
		record.UriStem,
		record.Status,
		record.UriQuery,
		record.EdgeResultType,
		record.EdgeRequestId,
		record.HostHeader,
		record.Protocol,
		record.BytesReceived,
		record.TimeTaken,
		record.Referer,
		record.UserAgent,
		record.Cookie,
		record.Time,
	}

	series.Points = append(series.Points, new_array)

	return &series
}

func AddToInflux(record *cdnlysis.LogRecord) {
	if cachedInfluxConn == nil {
		makeInfluxSession()
	}

	value := makeRecord(record)

	if err := cachedInfluxConn.WriteSeries([]*client.Series{value}); err != nil {
		log.Println("Cannot add to Influx", err)
		return
	}
}
