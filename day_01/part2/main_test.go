package main

import "testing"

func TestCalcFuel(t *testing.T) {
	fuel := calcFuel(12)
	if fuel != 2 {
		t.Errorf("Fuel incorrect, got: %d, want: %d", fuel, 2)
	}

	fuel = calcFuel(14)
	if fuel != 2 {
		t.Errorf("Fuel incorrect, got: %d, want: %d", fuel, 2)
	}

	fuel = calcFuel(1969)
	if fuel != 654 {
		t.Errorf("Fuel incorrect, got: %d, want: %d", fuel, 654)
	}

	fuel = calcFuel(100756)
	if fuel != 33583 {
		t.Errorf("Fuel incorrect, got: %d, want: %d", fuel, 33583)
	}
}

func TestCalcFuelRecursive(t *testing.T) {
	fuel := calcFuelRecursive(12)
	if fuel != 2 {
		t.Errorf("Fuel incorrect, got: %d, want: %d", fuel, 2)
	}

	fuel = calcFuelRecursive(1969)
	if fuel != 966 {
		t.Errorf("Fuel incorrect, got: %d, want: %d", fuel, 966)
	}

	fuel = calcFuelRecursive(100756)
	if fuel != 50346 {
		t.Errorf("Fuel incorrect, got: %d, want: %d", fuel, 50346)
	}
}
