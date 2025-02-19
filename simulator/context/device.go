package context

import (
	"errors"
	"fmt"
	"github.com/brocaar/lorawan"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type Device struct {
	DevId uint16
	// BanchMark  variables
	Packet_tx, Packet_rx uint8
	PacketLoss           int
	Durations            DurationSlice
	Start                time.Time
	TotalTime 			 float64
	FsmState             int
	Confirmed            bool
	//devEUI lorawan.EUI64
	// AppKey.
	//appKey lorawan.AES128Key
	// Application session-key.
	appSKey lorawan.AES128Key
	// Network session-key.
	nwkSKey lorawan.AES128Key
	// devAddr
	DevAddr lorawan.DevAddr
	// DevEUI.
	//devEUI lorawan.EUI64
	// FPort used for sending uplinks.
	fPort uint8
	// Uplink frame-counter.
	fCntUp uint32
	// Downlink frame-counter.
	fCntDown uint32
	// Payload (plaintext) which the device sends as uplink.
	payload []byte
	DownlinkHandleFunc func() error

	DoneRecv chan bool
}

func (d *Device) Marshall() ([]byte, bool){

	phyLoRa, ok := d.UplinkData()
	if !ok {
		log.Fatalf("Error Marshall Phy Lora frame ")
	}
	payload, err := phyLoRa.MarshalBinary()
	if err != nil {
		log.Fatalf("%v", errors.New("Error marshall binary  device Data"))
		return nil, false
	}
	return payload, true
}

func (device *Device) init( id uint16){
	device.DevId = id
	//device.DevAddr = counter.getAddr()
	device.DevAddr = DevIDtoHEx(id)
	device.DoneRecv = make(chan bool)
}

func (device *Device) GetDevID() (uint16) {
	//b := make([]byte, 2)
	//binary.BigEndian.PutUint16(b, device.DevId)
	//return hex.EncodeToString(b)
	return device.DevId
}

func (device *Device) SetMessagePayload( msg string ){
	device.payload = []byte(msg)
}
func (device *Device) ElapsedTime(){
	t := time.Now()
	elapsed := t.Sub(device.Start)
	device.TotalTime += elapsed.Seconds()
	device.Durations = append(device.Durations, elapsed)
}

// dataUp sends an data uplink.
func (d *Device) UplinkData() (lorawan.PHYPayload, bool) {

	mType := lorawan.UnconfirmedDataUp
	if d.Confirmed {
		mType = lorawan.ConfirmedDataUp
	}

	phy := lorawan.PHYPayload{
		MHDR: lorawan.MHDR{
			MType: mType,
			Major: lorawan.LoRaWANR1,
		},

		MACPayload: &lorawan.MACPayload{

			FHDR: lorawan.FHDR{
				DevAddr: d.DevAddr,
				FCnt:    d.fCntUp,
				FCtrl: lorawan.FCtrl{
					ADR: false,
				},
			},

			FPort: &d.fPort,

			FRMPayload: []lorawan.Payload{
				&lorawan.DataPayload{
					Bytes: d.payload,
				},
			},

		},
	}

	if err := phy.EncryptFRMPayload(d.appSKey); err != nil {
		logrus.WithError(err).Error("simulator: encrypt FRMPayload error")
		return  lorawan.PHYPayload{} , false
	}

	if err := phy.SetUplinkDataMIC(lorawan.LoRaWAN1_0, 0, 0, 0, d.nwkSKey, d.nwkSKey); err != nil {
		logrus.WithError(err).Error("simulator: set uplink data mic error")
		return lorawan.PHYPayload{} , false
	}

	d.fCntUp++

	return phy, true
}

func (d *Device)  UpLinkInfo() []string {
	return []string{
		fmt.Sprintf("%d", d.DevId),
		fmt.Sprintf("%s", "uplink"),
		fmt.Sprintf("%d", d.Packet_tx),
		fmt.Sprintf("%t", false),
		fmt.Sprintf("%s", d.Start.Format("15:04:05")),
	}
}

func (d *Device) DownLinkInfo(recv bool ) []string {
	return []string{
		fmt.Sprintf("%d", d.DevId),
		fmt.Sprintf("%s", "downlink"),
		fmt.Sprintf("%d", d.Packet_rx),
		fmt.Sprintf("%t", recv),
		fmt.Sprintf("%f", d.Durations[len(d.Durations) -1].Seconds()),
	}
}

func (d *Device)  UpLinkInfoResume() []string {
	return []string{
		fmt.Sprintf("%d", d.DevId),
		fmt.Sprintf("%s", "uplink"),
		fmt.Sprintf("%d", d.Packet_tx),
		fmt.Sprintf("%f",0.0),
	}
}

func (d *Device) DownLinkInfoResume() []string {
	return []string{
		fmt.Sprintf("%d", d.DevId),
		fmt.Sprintf("%s", "downlink"),
		fmt.Sprintf("%d", d.Packet_rx),
		fmt.Sprintf("%f", d.TotalTime),
	}
}

func CreateDevicesForSimulate(devicesLen int){
	if devicesLen < 1 {
		log.Fatalf("Number of devices is not valid")
	}
	for i := 0; i < devicesLen; i++ {
		DevicesContext_Self().NewDevice()
	}
}