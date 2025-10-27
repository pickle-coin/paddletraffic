<script lang="ts">
	// Pure, reusable map component using maplibre-gl
	import maplibregl from 'maplibre-gl';
	import 'maplibre-gl/dist/maplibre-gl.css';
	import { onMount } from 'svelte';
	import { initMarkerManager } from '$lib/map/MarkerManager.svelte';

	interface MapProps {
		children?: any;
	}

	let { children }: MapProps = $props();

	let mapContainer: HTMLDivElement;
	let map: maplibregl.Map;

	onMount(() => {
		// Initialize the map
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

		// Initialize marker manager (handles all marker reactivity)
		const cleanupMarkers = initMarkerManager(map);

		// Cleanup
		return () => {
			cleanupMarkers();
			map.remove();
		};
	});
</script>

<div class="relative m-0 h-screen w-full">
	<div bind:this={mapContainer} class="h-screen w-full"></div>

	{@render children?.()}
</div>
