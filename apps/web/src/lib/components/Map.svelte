<script lang="ts">
	// basic map component using maplibre-gl
	import maplibregl from 'maplibre-gl';
	import 'maplibre-gl/dist/maplibre-gl.css';
	import { onMount } from 'svelte';

	let mapContainer: HTMLDivElement;

	onMount(() => {
			const map = new maplibregl.Map({
				container: mapContainer, // container id
				style: 'https://demotiles.maplibre.org/style.json', // style URL
				center: [0, 0], // starting position [lng, lat]
				zoom: 1 // starting zoom
			});
	})

	let { children } = $props()
</script>

<div class="relative h-screen w-full m-0">
	<div bind:this={mapContainer} class="h-screen w-full" />
		
	{@render children?.()}
</div>