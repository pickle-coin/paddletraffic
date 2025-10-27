<script lang="ts">
	// main map page following figma design using map component and other features composed together
	import Map from '$lib/components/Map.svelte';
	import CourtDrawer from '$lib/components/CourtDrawer.svelte';
	import { courtsState } from '$lib/stores/courts.svelte';

	let drawerOpen = $state(false);

	// Open drawer when a court is selected, close when deselected
	$effect(() => {
		drawerOpen = courtsState.selectedCourtId !== null;
	});

	// When drawer is closed manually, deselect the marker
	$effect(() => {
		if (!drawerOpen && courtsState.selectedCourtId !== null) {
			courtsState.selectCourt(null);
		}
	});
</script>

<Map>
	<CourtDrawer bind:open={drawerOpen} court={courtsState.selectedCourt} />
</Map>
