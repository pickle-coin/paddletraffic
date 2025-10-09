<script lang="ts">
	import { onMount } from 'svelte';

	let leafletContainer: HTMLDivElement;
	let mapboxContainer: HTMLDivElement;

	onMount(async () => {
		// Define common center and zoom for both maps
		const center = { lat: 37.7749, lng: -122.4194 };
		const zoom = 11;

		// Import Leaflet only in browser
		const L = (await import('leaflet')).default;

		// Initialize Leaflet map
		const leafletMap = L.map(leafletContainer).setView([center.lat, center.lng], zoom);
		L.tileLayer('https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png', {
			attribution: '© OpenStreetMap contributors, © CARTO'
		}).addTo(leafletMap);
		L.marker([center.lat, center.lng]).addTo(leafletMap).bindPopup('San Francisco');

		// Import MapLibre GL JS only in browser
		const maplibregl = (await import('maplibre-gl')).default;

		// Initialize MapLibre GL JS map with matching style
		const maplibreMap = new maplibregl.Map({
			container: mapboxContainer,
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
			center: [center.lng, center.lat],
			zoom: zoom - 1 // MapLibre zoom is offset by 1 from Leaflet
		});
		new maplibregl.Marker()
			.setLngLat([center.lng, center.lat])
			.setPopup(new maplibregl.Popup().setText('San Francisco'))
			.addTo(maplibreMap);
	});
</script>

<div class="container mx-auto p-8">
	<h1 class="mb-8 text-3xl font-bold">Leaflet vs MapLibre GL JS Demo</h1>

	<div class="mb-12 grid grid-cols-1 gap-8 lg:grid-cols-2">
		<!-- Leaflet Demo -->
		<div>
			<h2 class="mb-4 text-2xl font-semibold">Leaflet</h2>
			<div bind:this={leafletContainer} class="mb-4 h-96 rounded border border-gray-300"></div>
			<div class="rounded bg-gray-100 p-4">
				<h3 class="mb-2 text-lg font-medium">Code:</h3>
				<pre class="overflow-x-auto text-sm"><code>{`import L from 'leaflet';

const map = L.map('map')
  .setView([37.7749, -122.4194], 11);

L.tileLayer(
  'https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png',
  { attribution: '© OpenStreetMap, © CARTO' }
).addTo(map);

L.marker([37.7749, -122.4194])
  .addTo(map)
  .bindPopup('San Francisco');`}</code></pre>
			</div>
		</div>

		<!-- MapLibre GL JS Demo -->
		<div>
			<h2 class="mb-4 text-2xl font-semibold">MapLibre GL JS</h2>
			<div bind:this={mapboxContainer} class="mb-4 h-96 rounded border border-gray-300"></div>
			<div class="rounded bg-gray-100 p-4">
				<h3 class="mb-2 text-lg font-medium">Code:</h3>
				<pre class="overflow-x-auto text-sm"><code>{`import maplibregl from 'maplibre-gl';

const map = new maplibregl.Map({
  container: 'map',
  style: {
    version: 8,
    sources: {
      'carto-light': {
        type: 'raster',
        tiles: [
          'https://a.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png'
        ],
        tileSize: 256,
        attribution: '© OpenStreetMap, © CARTO'
      }
    },
    layers: [{
      id: 'carto-light-layer',
      type: 'raster',
      source: 'carto-light'
    }]
  },
  center: [-122.4194, 37.7749],
  zoom: 10  // Offset by 1 from Leaflet for matching view
});

new maplibregl.Marker()
  .setLngLat([-122.4194, 37.7749])
  .setPopup(
    new maplibregl.Popup()
      .setText('San Francisco')
  )
  .addTo(map);`}</code></pre>
			</div>
		</div>
	</div>
</div>
