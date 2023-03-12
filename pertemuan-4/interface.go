package main

import "math"

type bola struct {
	jari   float64
	tinggi float64
}

type kerucut struct {
	selimut float64
	jari    float64
	tinggi  float64
}

type dimensi2 interface {
	luas() float64
	keliling() float64
}

type dimensi3 interface {
	volume() float64
}

type hitung interface {
	dimensi2
	dimensi3
}

func (b *bola) setUkuranBola(jari float64, tinggi float64) {
	b.jari = jari
	b.tinggi = tinggi
}

func (k *kerucut) setUkuranKerucut(selimut float64, jari float64, tinggi float64) {
	k.selimut = selimut
	k.jari = jari
	k.tinggi = tinggi
}

func (b *bola) luas() float64 {
	return 4 * math.Pi * math.Pow(b.jari, 2)
}

func (b *bola) keliling() float64 {
	return 4 / 3 * math.Pi * math.Pow(b.jari/2, 2)
}

func (b *bola) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(b.jari/2, 2) * b.tinggi
}

func (k *kerucut) luas() float64 {
	return math.Pi * k.jari * (k.jari + k.selimut)
}

func (k *kerucut) keliling() float64 {
	return 2 * math.Pi * k.jari
}

func (k *kerucut) volume() float64 {
	return 3 * math.Pi * math.Pow(k.jari/2, 2) * k.tinggi
}
