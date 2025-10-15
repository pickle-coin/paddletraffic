<script lang="ts">
	// main map page following figma design using map component and other features composed together
	import Map from '$lib/components/Map.svelte';
	import Popup from '$lib/components/Popup.svelte';

	let open = $state(true);
	const toggle = () => { open = !open };
</script>

<Map>
	<button onclick={toggle} class="absolute left-4 top-4 rounded-md bg-red-600 p-2 hover:cursor-pointer"> TEST TEXT </button>

	<Popup
		bind:open
		class="absolute bottom-0 left-0 right-0 w-full"
		title="Aaron's HQ"
		courts_occupied={2}
		groups_waiting={1}
		estimated_wait_time_minutes={60}
	></Popup>
</Map>
