package main

import "testing"

/**
author:Boyn
date:2020/2/13
*/

var tracks = []*Track{
	{"Go", "Delilah", "From the roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("4m28s")},
	{"Bial", "Bob", "As I am", 2006, length("2m59s")},
	{"Cat", "Alian", "Age 18", 2018, length("3m35s")},
	{"Go", "Carl", "1998", 1998, length("4m12s")},
}

func TestPrintTrack(t *testing.T) {
	PrintTracks(tracks)
}

func TestSortByYear(t *testing.T) {
	PrintTracks(tracks)
	SortByYear(tracks)
	PrintTracks(tracks)
}

func TestSortByArtist(t *testing.T) {
	PrintTracks(tracks)
	SortByArtist(tracks)
	PrintTracks(tracks)
}
