package adapters

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bchanona/profile_warmheart_backend/User/domain"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

// InitMQTT inicializa la conexión al broker MQTT
func InitMQTT() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://3.216.41.175:1883") // Dirección del broker MQTT
	opts.SetClientID("mqtt_publisher")
	opts.SetUsername("carlos")
	opts.SetPassword("carlos")

	// Crear cliente MQTT
	client = mqtt.NewClient(opts)

	// Intentar conectar
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al conectar con MQTT: %v", token.Error())
	}
	fmt.Println("Conectado al broker MQTT")
}

// PublishUserData publica el id_usuario e id_device en MQTT
func PublishUserData(idUser, idDevice int) error {
	// Verifica si el cliente MQTT está conectado
	if client == nil || !client.IsConnected() {
		log.Println("El cliente MQTT no está conectado.")
		return fmt.Errorf("cliente MQTT no conectado")
	}

	data := domain.DeviceData{
		IdUser:   idUser,
		IdDevice: idDevice,
	}

	// Convertir los datos a JSON
	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error al serializar los datos: %v", err)
		return fmt.Errorf("error al serializar los datos")
	}

	// Publicar en MQTT en ambos tópicos
	topics := []string{"device.data"}
	for _, topic := range topics {
		token := client.Publish(topic, 0, false, payload)
		token.Wait() // Esperar a que se complete la publicación

		if token.Error() != nil {
			log.Printf("Error al publicar el mensaje en %s: %v", topic, token.Error())
			return fmt.Errorf("error al publicar el mensaje en MQTT en el tópico %s", topic)
		}

		fmt.Printf("Mensaje enviado a MQTT en %s: %s\n", topic, string(payload))
	}

	return nil
}
