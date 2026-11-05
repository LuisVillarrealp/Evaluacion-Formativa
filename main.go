package main

import (
	"fmt"
	"time"
)

type Medicine struct{
	name	string
	manufacturer string
	manufacturerDate time.Time
	shelfLife int
}
func NuevaMedicine(name string, manufacturer string, manufacturerDate time.Time, shelfLife int)Medicine{
	return Medicine{
		name: name,
		manufacturer: manufacturer,
		manufacturerDate: manufacturerDate,
		shelfLife: shelfLife,
	}
}
func (m Medicine) Getname() string {
	return m.name
}
func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}
func (m Medicine) GetManufacturerDate() time.Time {
	return m.manufacturerDate
}
func (m Medicine) GetShelfLife() int{
	return m.shelfLife
}
func (m *Medicine) SetName(nuevoNombre string){
	m.name = nuevoNombre
}
func (m *Medicine) SetManufacturer(nuevoManufacturer string){
	m.manufacturer = nuevoManufacturer
}
func (m *Medicine) SetManufacturerDate(nuevoManufacturerDate time.Time){
	m.manufacturerDate = nuevoManufacturerDate
}
func (m *Medicine) SetShelfLife(nuevoShelfLife int){
	m.shelfLife = nuevoShelfLife
}

func (m Medicine) MostrarInformacionMedicine() {
	fmt.Println("Nombre:", m.Getname())
	fmt.Println("Fabricante:", m.GetManufacturer())
	fmt.Println("Fecha de fabricación:", m.GetManufacturerDate().Format("02/01/2006"))
	fmt.Println("Vida útil (meses):", m.GetShelfLife())
}



type Tablet struct{
	Information Medicine
	dosePerTablet int
	isPrescription bool
}
type Syrup struct{
	Information Medicine
	volume int
	flavor string
}
func NewTablet(NewMedicine Medicine, dosePerTablet int, isPrescription bool) Tablet {
	return Tablet{
		Information: NewMedicine,
		dosePerTablet: dosePerTablet,
		isPrescription: isPrescription,
	}
}
func NewSyrup(NewMedicine Medicine, volume int, flavor string) Syrup {
	return Syrup{
		Information: NewMedicine,
		volume: volume,
		flavor: flavor,
	}
}
func (t Tablet) GetDosePerTablet() int {
	return t.dosePerTablet
}
func (t Tablet) GetIsPrescription() bool {
	return t.isPrescription
}
func (t Tablet) GetInformation() Medicine {
	return t.Information
}
func (t Tablet)MostrarDetalles(){
	fmt.Println("Tableta")
	med := t.GetInformation()        
	med.MostrarInformacionMedicine()
	fmt.Println("Dosis por tableta (mg):", t.GetDosePerTablet())
	fmt.Println("Requiere receta:", t.GetIsPrescription(),)
	fmt.Println()
}
func (s Syrup) Getvolume() int {
	return s.volume
}
func (s Syrup) Getflavor() string {
	return s.flavor
}
func (s Syrup) GetInformation() Medicine {
	return s.Information
}
func (s Syrup)MostrarDetalles(){
	fmt.Println("Syrup")
	med := s.GetInformation()       
	med.MostrarInformacionMedicine()
	fmt.Println("Volumen (ml):", s.Getvolume())
	fmt.Println("Sabor:", s.Getflavor(),)
	fmt.Println()
}



func (m Medicine) CalcularFecha()time.Time{
	return m.manufacturerDate.AddDate(0, m.shelfLife, 0)
}
func main(){
	tabletas := []Tablet{
		NewTablet(NuevaMedicine("Paracetamol", "ParacetamolOwner", time.Date(2026, time.March, 15, 0, 0, 0, 0, time.UTC), 24),500, false,),
		NewTablet(NuevaMedicine("Amoxicilina", "AmoxilinaOwner", time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC), 18),875, true,),
	}

	syrups := []Syrup{
		NewSyrup(NuevaMedicine("Ambroxol", "AmbroxolCompany", time.Date(2026, time.January, 10, 0, 0, 0, 0, time.UTC), 12),120, "Mora",),
		NewSyrup(NuevaMedicine("Loratadina", "LoratadinaCompany", time.Date(2026, time.June, 20, 0, 0, 0, 0, time.UTC), 24),60, "Naranja",),
	}
	for _, t := range tabletas {
		t.MostrarDetalles()
	}
	for _, s := range syrups {
		s.MostrarDetalles()
	}

	fmt.Println("Antes de Cambiar",tabletas[1].Information.GetManufacturer())
	tabletas[1].Information.SetManufacturer("AmoxilinaCompany")
	fmt.Println("Despues de Cambiar",tabletas[1].Information.GetManufacturer())


	for _, t := range tabletas {
		med := t.GetInformation()
		fmt.Println(med.Getname(), "caduca el", med.CalcularFecha().Format("02/01/2006"))
	}
	for _, s := range syrups {
		med := s.GetInformation()
		fmt.Println(med.Getname(), "caduca el", med.CalcularFecha().Format("02/01/2006"))
	}
}