package dashboard

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// combines config and reflector data into json for the dashboard view
func (d *Dashboard) showMetadata(c *gin.Context) {
	d.Reflector.refreshIfNeeded()
	d.Reflector.Lock.Lock()
	defer d.Reflector.Lock.Unlock()
	c.JSON(200, gin.H{
		"sysop_email":        d.Config.Email,
		"ipV4":               d.Config.IPv4,
		"ipV6":               d.Config.IPv6,
		"reflector_callsign": d.Reflector.ReflectorData.Callsign,
		"reflector_version":  d.Reflector.ReflectorData.Version,
		"dashboard_version":  d.Version,
	})
}

// showStatus returns runtime information about the reflector and dashboard.
func (d *Dashboard) showStatus(c *gin.Context) {
	r := d.Reflector
	r.refreshIfNeeded()
	status := struct {
		LastUpdateUnixTime     int64  `json:"lastupdate"`
		LastDataUpdateUnixTime int64  `json:"lasturfdupdate"`
		ReflectorStatus        string `json:"reflectorstatus"`
		ReflectorUptimeSeconds int64  `json:"reflectoruptimeseconds"`
	}{
		LastUpdateUnixTime:     r.LastUpdateCheckTime.Unix(),
		LastDataUpdateUnixTime: r.ReflectorData.FileTime.Unix(),
		ReflectorStatus:        r.Status,
		ReflectorUptimeSeconds: r.UptimeSeconds,
	}

	c.JSON(200, status)
}

// showCountries returns the country callsign data in JSON.
func (d *Dashboard) showCountries(c *gin.Context) {
	c.JSON(200, d.CountryCallsigns)
}

// showReflectorJSON returns the entire reflector stuct in JSON.
func (d *Dashboard) showReflectorJSON(c *gin.Context) {
	d.Reflector.Lock.Lock()
	defer d.Reflector.Lock.Unlock()
	c.JSON(200, d.Reflector)
}

func (d *Dashboard) showStationDataJSON(c *gin.Context) {
	d.Reflector.refreshIfNeeded()
	d.Reflector.Lock.Lock()
	defer d.Reflector.Lock.Unlock()
	type stationData struct {
		Callsign       string  `json:"callsign"`
		CallsignSuffix string  `json:"callsignsuffix"`
		ViaNode        string  `json:"vianode"`
		OnModule       string  `json:"onmodule"`
		LastHeard      string  `json:"lastheard"`
		Country        Country `json:"country"`
	}
	var data struct {
		Stations []stationData `json:"stations"`
	}
	for i, station := range d.Reflector.ReflectorData.Stations {
		callsignSplit := strings.Fields(station.Callsign)
		if len(callsignSplit) < 2 {
			callsignSplit = append(callsignSplit, " ")
		}
		country, err := d.CountryCallsigns.GetCountryFromCallsign(callsignSplit[0])
		if err != nil {
			country = Country{}
		}
		data.Stations = append(data.Stations, stationData{
			Callsign:       callsignSplit[0],
			CallsignSuffix: callsignSplit[1],
			ViaNode:        station.ViaNode,
			OnModule:       station.OnModule,
			LastHeard:      station.LastHeardTime,
			Country:        country,
		})
		if i >= d.Config.LastHeard-1 {
			break
		} // respect the last heard
	}
	c.JSON(200, data)
}

func (d *Dashboard) showPeers(c *gin.Context) {
	d.Reflector.refreshIfNeeded()
	d.Reflector.Lock.Lock()
	defer d.Reflector.Lock.Unlock()
	c.JSON(200, d.Reflector.ReflectorData.Peers)
}

func (d *Dashboard) showLinksDataJSON(c *gin.Context) {
	d.Reflector.refreshIfNeeded()
	c.JSON(200, d.Reflector.ReflectorData.Nodes)
}

func (d *Dashboard) showModulesInUseJSON(c *gin.Context) {
	type moduleInfo struct {
		Name      string   `json:"name"`
		Callsigns []string `json:"callsigns"`
	}

	modules := d.Reflector.GetModules()
	// Go Maps are not in any order.. binary trees.. make order from the chaos!

	// make a list of keys.
	keys := make([]string, 0, len(modules))
	for k := range modules {
		keys = append(keys, k)
	}

	// sort the keys
	sort.Strings(keys)

	// make moduleData by iterating the sorted keys
	moduleData := make([]moduleInfo, 0, len(keys))
	for _, name := range keys {
		//sort the callsigns
		callSigns := modules[name]
		sort.Strings(callSigns)

		moduleData = append(moduleData, moduleInfo{Name: name, Callsigns: callSigns})
	}
	c.JSON(200, moduleData)
}
