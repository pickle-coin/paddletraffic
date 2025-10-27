<script lang="ts">
	// Pure, reusable map component using maplibre-gl
	import maplibregl from 'maplibre-gl';
	import 'maplibre-gl/dist/maplibre-gl.css';
	import { onMount } from 'svelte';
	import type { Court } from '$lib/data/dummyCourts';

	interface MapProps {
		courts: Court[];
		selectedCourtId: number | null;
		onMarkerClick: (courtId: number) => void;
		children?: any;
	}

	let { courts, selectedCourtId, onMarkerClick, children }: MapProps = $props();

	let mapContainer: HTMLDivElement;
	let markers: Map<number, { marker: maplibregl.Marker; element: HTMLElement }> = new Map();

	onMount(() => {
		const map = new maplibregl.Map({
			container: mapContainer,
			style: {
				version: 8,
				sources: {
					'carto-light': {
						type: 'raster',
						tiles: [
							'https://a.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png',
							'https://b.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png',
							'https://c.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png'
						],
						tileSize: 256,
						attribution: '© OpenStreetMap contributors, © CARTO'
					}
				},
				layers: [
					{
						id: 'carto-light-layer',
						type: 'raster',
						source: 'carto-light',
						minzoom: 0,
						maxzoom: 22
					}
				]
			},
			center: [-111.891, 40.7608], // Salt Lake City [lng, lat]
			zoom: 12 // starting zoom
		});

		// Add markers for each court
		courts.forEach((court) => {
			// TODO: Future enhancement - color markers based on busyness (green=empty, yellow=some waiting, red=busy)
			// For now, all markers are green since dummy data has 0 courts occupied
			const marker = new maplibregl.Marker({ color: '#22c55e' })
				.setLngLat([court.location.coordinates.lon, court.location.coordinates.lat])
				.addTo(map);

			// Get the marker's DOM element to add click handler and styling
			const el = marker.getElement();
			el.style.cursor = 'pointer';

			// Get the SVG element inside the marker for styling
			const svg = el.querySelector('svg');
			if (svg) {
				svg.style.transition = 'stroke 0.2s ease';
				svg.style.overflow = 'visible'; // Prevent outline from being clipped
			}

			el.addEventListener('click', () => {
				onMarkerClick(court.id);
			});

			markers.set(court.id, { marker, element: el });
		});
	});

	// Update marker styles when selection changes
	$effect(() => {
		markers.forEach(({ element }, courtId) => {
			const svg = element.querySelector('svg');
			if (svg) {
				const path = svg.querySelector('path');
				if (path) {
					if (courtId === selectedCourtId) {
						// Blue outline effect for selected marker
						path.setAttribute('stroke', '#3b82f6');
						path.setAttribute('stroke-width', '3');
					} else {
						// No outline for unselected
						path.setAttribute('stroke', 'none');
						path.setAttribute('stroke-width', '0');
					}
				}
			}
		});
	});
</script>

<div class="relative m-0 h-screen w-full">
	<div bind:this={mapContainer} class="h-screen w-full"></div>

	{@render children?.()}
</div>
