<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import 'leaflet/dist/leaflet.css'; // required by leaflet, may be only needed on the +page.svelte?

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}
