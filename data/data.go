package data

type GPUInfo struct {
	Address    string     `json:"address"`
	ModelInfo  ModelInfo  `json:"modelinfo"`
	UsageStats UsageStats `json:"usagestats"`
}

type ModelInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type UsageStats struct {
	Tmp string `json:"tmp"`
}

var MyGPUInfo GPUInfo = GPUInfo{}

// API 서버 address
var HttpAddress string = "localhost:8080"

// 프로메테우스 서버 address
var PrometheusAddress string = "http://tritonserver:8002/metrics"

//"ssl.ws.ahri.world:8002"

// Manager address
var ManagerAddress *string
