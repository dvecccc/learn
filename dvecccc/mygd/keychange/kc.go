package keychange

import (
	"errors"
	"fmt"
	"github.com/dvecccc/mygd/rabbitmq"
	"log"
)

type MicroInfo struct {
	Ip       string
	Port     string
	Key      string
	Exchange string
}

type MicroService struct {
	Ip               string
	Port             string
	ExchangesAndKeys map[string][]string
	rabbitMQs        map[string]*rabbitmq.RabbitMQ
}

func (s *MicroService) logInfo(message string) {
	log.Printf("[ERROR] %s\n", message)
}

func join(exchange string, key string) string {
	return exchange + key
}

func (s *MicroService) AddExchangeAndKey(exchange string, key string) error {
	mqTopic := rabbitmq.NewRabbitMQTopic(exchange, key)
	if mqTopic == nil {
		s.logInfo(fmt.Sprintf("rabbitmq: [exchange: %s, key: %s] create failed.", exchange, key))
		return errors.New(fmt.Sprintf("rabbitmq: [exchange: %s, key: %s] create failed.", exchange, key))
	}
	keySlice := s.ExchangesAndKeys[exchange]
	keySlice = append(keySlice, key)
	s.ExchangesAndKeys[exchange] = keySlice
	s.rabbitMQs[join(exchange, key)] = mqTopic
	s.logInfo(fmt.Sprintf("rabbitmq: [exchange: %s, key: %s] create successs.", exchange, key))
	return nil
}

func (s *MicroService) OnlyChangeExchange(oldExchange, newExchange string) error {
	keys := s.ExchangesAndKeys[oldExchange]
	newKeys := s.ExchangesAndKeys[newExchange]
	if newKeys == nil {
		s.ExchangesAndKeys[newExchange] = keys
	} else {
		newKeys = append(newKeys, keys...)
		s.ExchangesAndKeys[newExchange] = newKeys
	}
	for _, key := range keys {
		delete(s.rabbitMQs, join(oldExchange, key))
		mqTopic := rabbitmq.NewRabbitMQTopic(newExchange, key)
		s.rabbitMQs[join(newExchange, key)] = mqTopic
	}
	return nil
}

func (s *MicroService) OnlyChangeKeys(exchange string, keys []string) error {
	oldKeys := s.ExchangesAndKeys[exchange]
	for _, key := range oldKeys {
		delete(s.rabbitMQs, join(exchange, key))
	}
	s.ExchangesAndKeys[exchange] = keys
	for _, key := range keys {
		mqTopic := rabbitmq.NewRabbitMQTopic(exchange, key)
		s.rabbitMQs[join(exchange, key)] = mqTopic
	}
	return nil
}

func (s *MicroService) OnlyDeleteKey(exchange string, key string) error {
	keys := s.ExchangesAndKeys[exchange]
	for idx, k := range keys {
		if k == key {
			keys = append(keys[:idx], keys[idx+1:]...)
			s.ExchangesAndKeys[exchange] = keys
			break
		}
	}
	delete(s.rabbitMQs, join(exchange, key))
	return nil
}

func (s *MicroService) DeleteExchange(exchange string) error {
	keys := s.ExchangesAndKeys[exchange]
	for _, key := range keys {
		delete(s.rabbitMQs, join(exchange, key))
	}
	delete(s.ExchangesAndKeys, exchange)
	return nil
}

func ChangeKey(info *MicroInfo) {

}

func connectionContainer(ip string, port string) {

}

func Add() {

}