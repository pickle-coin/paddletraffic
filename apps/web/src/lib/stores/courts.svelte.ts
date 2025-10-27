// Universal reactivity with Svelte 5 runes
// Single source of truth for courts data and selection state
import type { Court } from '$lib/data/dummyCourts';
import { dummyCourts } from '$lib/data/dummyCourts';

// Reactive state using runes
let courts = $state<Court[]>(dummyCourts);
let selectedCourtId = $state<number | null>(null);

// Derived state for selected court
const selectedCourt = $derived(
	selectedCourtId ? courts.find((c) => c.id === selectedCourtId) ?? null : null
);

export const courtsState = {
	// Getters for reactive values
	get courts() {
		return courts;
	},
	get selectedCourtId() {
		return selectedCourtId;
	},
	get selectedCourt() {
		return selectedCourt;
	},

	// Actions
	selectCourt(id: number | null) {
		selectedCourtId = id;
	},

	toggleCourt(id: number) {
		if (selectedCourtId === id) {
			selectedCourtId = null;
		} else {
			selectedCourtId = id;
		}
	},

	// Future: Add methods for API calls
	// async fetchCourts() { ... }
	// async updateCourt(court: Court) { ... }
};
