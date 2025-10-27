// Reactive marker manager that handles all marker operations and state
// Combines marker utilities with reactive orchestration using Svelte 5 runes
import maplibregl from 'maplibre-gl';
import type { Court } from '$lib/data/dummyCourts';
import { courtsState } from '$lib/stores/courts.svelte';

export interface MarkerData {
	marker: maplibregl.Marker;
	element: HTMLElement;
	path: SVGPathElement | null;
}

// Internal state
let markerMap = $state<Map<number, MarkerData>>(new Map());
let previouslySelectedId = $state<number | null>(null);
let mapInstance: maplibregl.Map | null = null;

/**
 * Create a marker for a court and add it to the map
 */
function createMarker(court: Court, map: maplibregl.Map): MarkerData {
	// TODO: Color markers based on busyness (green --> red)
	const marker = new maplibregl.Marker({ color: '#22c55e' })
		.setLngLat([court.location.coordinates.lon, court.location.coordinates.lat])
		.addTo(map);

	const element = marker.getElement();
	element.style.cursor = 'pointer';
	const svg = element.querySelector('svg');
	const path = svg?.querySelector('path') ?? null;

	if (svg) {
		svg.style.transition = 'stroke 0.2s ease';
		svg.style.overflow = 'visible'; // Prevent outline from being clipped
	}

	element.addEventListener('click', () => {
		courtsState.toggleCourt(court.id);
	});

	return { marker, element, path };
}

/**
 * Remove a specific marker from the map
 */
function removeMarker(markerData: MarkerData): void {
	markerData.marker.remove();
}

/**
 * Apply selection styling to a marker
 */
function selectMarker(markerData: MarkerData): void {
	if (markerData.path) {
		markerData.path.setAttribute('stroke', '#3b82f6'); // Blue outline
		markerData.path.setAttribute('stroke-width', '3');
	}
}

/**
 * Remove selection styling from a marker
 */
function deselectMarker(markerData: MarkerData): void {
	if (markerData.path) {
		markerData.path.setAttribute('stroke', 'none');
		markerData.path.setAttribute('stroke-width', '0');
	}
}

/**
 * Clear all markers from the map
 */
function clearAllMarkers(): void {
	markerMap.forEach((markerData) => removeMarker(markerData));
	markerMap.clear();
}

/**
 * Sync markers with current court data
 */
function syncMarkers(): void {
	if (!mapInstance) return;

	// Remove all existing markers
	clearAllMarkers();

	// Add markers for each court
	courtsState.courts.forEach((court) => {
		const markerData = createMarker(court, mapInstance!);
		markerMap.set(court.id, markerData);
	});
}

/**
 * Update marker selection styling
 */
function updateMarkerSelection(): void {
	// Deselect previous marker
	if (previouslySelectedId !== null) {
		const prevMarker = markerMap.get(previouslySelectedId);
		if (prevMarker) {
			deselectMarker(prevMarker);
		}
	}

	// Select new marker
	if (courtsState.selectedCourtId !== null) {
		const newMarker = markerMap.get(courtsState.selectedCourtId);
		if (newMarker) {
			selectMarker(newMarker);
		}
	}

	previouslySelectedId = courtsState.selectedCourtId;
}

/**
 * Initialize the marker manager with a map instance
 * Returns a cleanup function to be called on component unmount
 */
export function initMarkerManager(map: maplibregl.Map): () => void {
	mapInstance = map;

	// Reactively sync markers when courts change
	$effect(() => {
		syncMarkers();
	});

	// Reactively update selection when selectedCourtId changes
	$effect(() => {
		updateMarkerSelection();
	});

	// Return cleanup function
	return () => {
		clearAllMarkers();
		mapInstance = null;
	};
}
