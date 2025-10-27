<script lang="ts">
	// Pure, reusable map component using maplibre-gl
	import maplibregl from 'maplibre-gl';
	import 'maplibre-gl/dist/maplibre-gl.css';
	import { onMount } from 'svelte';
	import type { Court } from '$lib/data/dummyCourts';
	import * as markers from '$lib/map/marker-utils.svelte';

	interface MapProps {
		courts: Court[];
		selectedCourtId: number | null;
		onMarkerClick: (courtId: number) => void;
		children?: any;
	}

	interface MarkerData {
		marker: maplibregl.Marker;
		element: HTMLElement;
		path: SVGPathElement | null;
	}

	let { courts, selectedCourtId, onMarkerClick, children }: MapProps = $props();

	let mapContainer: HTMLDivElement;
	let map: maplibregl.Map;
	let markerMap = new Map<number, MarkerData>();
	let previouslySelectedId: number | null = null;

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
			markers.clearAllMarkers(markerMap);
			map.remove();
		};
	});

	// Reactively update markers when courts prop changes
	$effect(() => {
		if (!map) return;
		markers.syncMarkers(markerMap, courts, map, onMarkerClick);
	});

	// Update marker selection
	$effect(() => {
		markers.updateMarkerSelection(markerMap, selectedCourtId, previouslySelectedId);
	});
</script>

<div class="relative m-0 h-screen w-full">
	<div bind:this={mapContainer} class="h-screen w-full"></div>

	{@render children?.()}
</div>
