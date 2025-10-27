import maplibregl from 'maplibre-gl';
import type { Court } from '$lib/data/dummyCourts';

interface MarkerData {
	marker: maplibregl.Marker;
	element: HTMLElement;
	path: SVGPathElement | null;
}

/**
 * Create a marker for a court and add it to the map
 */
export function createMarker(
	court: Court,
	map: maplibregl.Map,
	onClick: (courtId: number) => void
): MarkerData {
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
		onClick(court.id);
	});

	return { marker, element, path };
}

/**
 * Remove a specific marker from the map
 */
export function removeMarker(markerData: MarkerData): void {
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
 * Update marker selection
 * Deselects the previous marker and selects the new one
 */
export function updateMarkerSelection(
	markers: Map<number, MarkerData>,
	selectedCourtId: number | null,
	previouslySelectedId: number | null
): number | null {
	// Deselect previous marker
	if (previouslySelectedId !== null) {
		const prevMarker = markers.get(previouslySelectedId);
		if (prevMarker) {
			deselectMarker(prevMarker);
		}
	}

	// Select new marker
	if (selectedCourtId !== null) {
		const newMarker = markers.get(selectedCourtId);
		if (newMarker) {
			selectMarker(newMarker);
		}
	}

	return selectedCourtId;
}

/**
 * Clear all markers from the map
 */
export function clearAllMarkers(markers: Map<number, MarkerData>): void {
	markers.forEach((markerData) => removeMarker(markerData));
	markers.clear();
}

/**
 * Sync markers with new court data
 */
export function syncMarkers(
	markers: Map<number, MarkerData>,
	courts: Court[],
	map: maplibregl.Map,
	onClick: (courtId: number) => void
): void {
	// Remove all existing markers
	clearAllMarkers(markers);

	// Add markers for each court
	courts.forEach((court) => {
		const markerData = createMarker(court, map, onClick);
		markers.set(court.id, markerData);
	});
}
