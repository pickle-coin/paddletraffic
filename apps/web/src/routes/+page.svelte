<script lang="ts">
	// main map page following figma design using map component and other features composed together
	import Map from '$lib/components/Map.svelte';
	import CourtDrawer from '$lib/components/CourtDrawer.svelte';
	import { dummyCourts } from '$lib/data/dummyCourts';

	let selectedCourtId = $state<number | null>(null);
	let drawerOpen = $state(false);

	// Find the selected court
	const selectedCourt = $derived(
		selectedCourtId ? dummyCourts.find((c) => c.id === selectedCourtId) ?? null : null
	);

	const handleMarkerClick = (courtId: number) => {
		if (selectedCourtId === courtId) {
			// Clicking the same marker - deselect and close drawer
			selectedCourtId = null;
			drawerOpen = false;
		} else {
			// Clicking a different marker - select it and open drawer
			selectedCourtId = courtId;
			drawerOpen = true;
		}
	};

	// When drawer is closed manually, also deselect the marker
	$effect(() => {
		if (!drawerOpen && selectedCourtId !== null) {
			selectedCourtId = null;
		}
	});
</script>

<Map courts={dummyCourts} {selectedCourtId} onMarkerClick={handleMarkerClick}>
	<CourtDrawer bind:open={drawerOpen} court={selectedCourt} />
</Map>
