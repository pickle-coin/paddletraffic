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
	let markers: Map<number, { marker: maplibregl.Marker; element: HTMLElement; path: SVGPathElement | null }> = new Map();
	let map: maplibregl.Map;

	onMount(() => {
		map = new maplibregl.Map({
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
			zoom: 12, // starting zoom
			doubleClickZoom: false // Disable double-tap/double-click zoom
		});

		// Cleanup
		return () => {
			markers.forEach(({ marker }) => marker.remove());
			markers.clear();
			map.remove();
		};
	});

	// Update markers when courts prop changes
	$effect(() => {
		if (!map) return;

		markers.forEach(({ marker }) => marker.remove());
		markers.clear();

		// Add markers for each court
		courts.forEach((court) => {
			// TODO: Color markers based on busyness (green=empty, yellow=some waiting, red=busy)
			const marker = new maplibregl.Marker({ color: '#22c55e' })
				.setLngLat([court.location.coordinates.lon, court.location.coordinates.lat])
				.addTo(map);

			// Get the marker's DOM element to add click handler and styling
			const el = marker.getElement();
			el.style.cursor = 'pointer';

			const svg = el.querySelector('svg');
			const path = svg?.querySelector('path') ?? null;

			if (svg) {
				svg.style.transition = 'stroke 0.2s ease';
				svg.style.overflow = 'visible'; // Prevent outline from being clipped
			}

			el.addEventListener('click', () => {
				onMarkerClick(court.id);
			});

			markers.set(court.id, { marker, element: el, path });
		});
	});

	// Update marker styles when selection changes
	$effect(() => {
		markers.forEach(({ path }, courtId) => {
			if (path) {
				if (courtId === selectedCourtId) {
					// Blue selected marker outline
					path.setAttribute('stroke', '#3b82f6');
					path.setAttribute('stroke-width', '3');
				} else {
					path.setAttribute('stroke', 'none');
					path.setAttribute('stroke-width', '0');
				}
			}
		});
	});
</script>

<div class="relative m-0 h-screen w-full">
	<div bind:this={mapContainer} class="h-screen w-full"></div>

	{@render children?.()}
</div>
