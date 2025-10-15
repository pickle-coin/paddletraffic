<script lang="ts">
	// main map page following figma design using map component and other features composed together
	import Map from '$lib/components/Map.svelte';
	import Popup from '$lib/components/Popup.svelte';

	let open = $state(true);
	const toggle = () => { open = !open };
</script>

<Map>
	<button onclick={toggle} class="absolute left-4 top-4 rounded-md bg-red-600 p-2 hover:cursor-pointer"> AARON'S HQ </button>

	<Popup
		bind:open
		title="Aaron's HQ"
		courts_occupied={2}
		total_courts={3}
		groups_waiting={1}
		estimated_wait_time_minutes={60}
	></Popup>
</Map>
