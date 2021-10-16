package main

type NSQTopic struct {
	Node         string `json:"node"`
	Hostname     string `json:"hostname"`
	TopicName    string `json:"topic_name"`
	Depth        int    `json:"depth"`
	MemoryDepth  int    `json:"memory_depth"`
	BackendDepth int    `json:"backend_depth"`
	MessageCount int    `json:"message_count"`
	Nodes        []struct {
		Node         string      `json:"node"`
		Hostname     string      `json:"hostname"`
		TopicName    string      `json:"topic_name"`
		Depth        int         `json:"depth"`
		MemoryDepth  int         `json:"memory_depth"`
		BackendDepth int         `json:"backend_depth"`
		MessageCount int         `json:"message_count"`
		Nodes        interface{} `json:"nodes"`
		Channels     []struct {
			Node                 string      `json:"node"`
			Hostname             string      `json:"hostname"`
			TopicName            string      `json:"topic_name"`
			ChannelName          string      `json:"channel_name"`
			Depth                int         `json:"depth"`
			MemoryDepth          int         `json:"memory_depth"`
			BackendDepth         int         `json:"backend_depth"`
			InFlightCount        int         `json:"in_flight_count"`
			DeferredCount        int         `json:"deferred_count"`
			RequeueCount         int         `json:"requeue_count"`
			TimeoutCount         int         `json:"timeout_count"`
			MessageCount         int         `json:"message_count"`
			ClientCount          int         `json:"client_count"`
			Nodes                interface{} `json:"nodes"`
			Clients              interface{} `json:"clients"`
			Paused               bool        `json:"paused"`
			E2EProcessingLatency struct {
				Count       int         `json:"count"`
				Percentiles interface{} `json:"percentiles"`
				Topic       string      `json:"topic"`
				Channel     string      `json:"channel"`
				Host        string      `json:"host"`
			} `json:"e2e_processing_latency"`
		} `json:"channels"`
		Paused               bool `json:"paused"`
		E2EProcessingLatency struct {
			Count       int         `json:"count"`
			Percentiles interface{} `json:"percentiles"`
			Topic       string      `json:"topic"`
			Channel     string      `json:"channel"`
			Host        string      `json:"host"`
		} `json:"e2e_processing_latency"`
	} `json:"nodes"`
	Channels []struct {
		Node                 string      `json:"node"`
		Hostname             string      `json:"hostname"`
		TopicName            string      `json:"topic_name"`
		ChannelName          string      `json:"channel_name"`
		Depth                int         `json:"depth"`
		MemoryDepth          int         `json:"memory_depth"`
		BackendDepth         int         `json:"backend_depth"`
		InFlightCount        int         `json:"in_flight_count"`
		DeferredCount        int         `json:"deferred_count"`
		RequeueCount         int         `json:"requeue_count"`
		TimeoutCount         int         `json:"timeout_count"`
		MessageCount         int         `json:"message_count"`
		ClientCount          int         `json:"client_count"`
		Nodes                interface{} `json:"nodes"`
		Clients              interface{} `json:"clients"`
		Paused               bool        `json:"paused"`
		E2EProcessingLatency struct {
			Count       int         `json:"count"`
			Percentiles interface{} `json:"percentiles"`
			Topic       string      `json:"topic"`
			Channel     string      `json:"channel"`
			Host        string      `json:"host"`
		} `json:"e2e_processing_latency"`
	} `json:"channels"`
	Paused               bool `json:"paused"`
	E2EProcessingLatency struct {
		Count       int         `json:"count"`
		Percentiles interface{} `json:"percentiles"`
		Topic       string      `json:"topic"`
		Channel     string      `json:"channel"`
		Host        string      `json:"host"`
	} `json:"e2e_processing_latency"`
	Message string `json:"message"`
}

type NSQChannel struct {
	Node          string `json:"node"`
	Hostname      string `json:"hostname"`
	TopicName     string `json:"topic_name"`
	ChannelName   string `json:"channel_name"`
	Depth         int    `json:"depth"`
	MemoryDepth   int    `json:"memory_depth"`
	BackendDepth  int    `json:"backend_depth"`
	InFlightCount int    `json:"in_flight_count"`
	DeferredCount int    `json:"deferred_count"`
	RequeueCount  int    `json:"requeue_count"`
	TimeoutCount  int    `json:"timeout_count"`
	MessageCount  int    `json:"message_count"`
	ClientCount   int    `json:"client_count"`
	Nodes         []struct {
		Node                 string        `json:"node"`
		Hostname             string        `json:"hostname"`
		TopicName            string        `json:"topic_name"`
		ChannelName          string        `json:"channel_name"`
		Depth                int           `json:"depth"`
		MemoryDepth          int           `json:"memory_depth"`
		BackendDepth         int           `json:"backend_depth"`
		InFlightCount        int           `json:"in_flight_count"`
		DeferredCount        int           `json:"deferred_count"`
		RequeueCount         int           `json:"requeue_count"`
		TimeoutCount         int           `json:"timeout_count"`
		MessageCount         int           `json:"message_count"`
		ClientCount          int           `json:"client_count"`
		Nodes                interface{}   `json:"nodes"`
		Clients              []interface{} `json:"clients"`
		Paused               bool          `json:"paused"`
		E2EProcessingLatency struct {
			Count       int         `json:"count"`
			Percentiles interface{} `json:"percentiles"`
			Topic       string      `json:"topic"`
			Channel     string      `json:"channel"`
			Host        string      `json:"host"`
		} `json:"e2e_processing_latency"`
	} `json:"nodes"`
	Clients              interface{} `json:"clients"`
	Paused               bool        `json:"paused"`
	E2EProcessingLatency struct {
		Count       int         `json:"count"`
		Percentiles interface{} `json:"percentiles"`
		Topic       string      `json:"topic"`
		Channel     string      `json:"channel"`
		Host        string      `json:"host"`
	} `json:"e2e_processing_latency"`
	Message string `json:"message"`
}

type NSQResponse struct {
	Message string `json:"message"`
}
