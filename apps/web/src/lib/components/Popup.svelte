<script lang="ts">
	import * as Drawer from '$lib/components/ui/drawer/index';
	import type { ClassValue } from 'clsx';
	import { X } from 'lucide-svelte';
	interface PopupProps {
		class?: ClassValue;
		open: boolean;
		title: string;
		courts_occupied: number;
		total_courts: number;
		groups_waiting: number;
		estimated_wait_time_minutes: number;
	}

	let {
		open = $bindable(),
		title,
		courts_occupied,
        total_courts,
		groups_waiting,
		estimated_wait_time_minutes
	}: PopupProps = $props();
</script>

<Drawer.Root bind:open modal={false}>
    <Drawer.Overlay class="-z-10 pointer-events-none"/>
	<Drawer.Content class="h-fit max-h-screen">
		<!-- CSS Gymnastics over here -->
		<div class="relative m-0 h-0 w-full p-0">
			<Drawer.Close>
				<X class="rounded-4xl absolute right-2 top-0 bg-neutral-200 p-1 text-neutral-600 hover:cursor-pointer" />
			</Drawer.Close>
		</div>

		<Drawer.Header>
			<div>
				<Drawer.Title class="text-2xl">{title}</Drawer.Title>
				<Drawer.Description class="text-sm">4.3 miles away</Drawer.Description>
			</div>
		</Drawer.Header>

        <!-- Main Content -->
        <div class="pb-4 flex flex-col">
            <div class="flex flex-col gap-2 p-1 px-2 pb-0">
                <!-- Row -->
                <div class="bar">
                    <div class="h-8 w-1.5 rounded-full bg-orange-400"></div>
                    <div class="flex w-full justify-between text-sm font-medium">
                        <span>Courts Occupied:</span>
                        <span>{courts_occupied}/{total_courts}</span>
                    </div>
                </div>
    
                <!-- Row -->
                <div class="bar">
                    <div class="h-8 w-1.5 rounded-full bg-green-400"></div>
                    <div class="flex w-full justify-between text-sm font-medium">
                        <span>Groups Waiting:</span>
                        <span>{groups_waiting}</span>
                    </div>
                </div>
    
                <!-- Row -->
                <div class="bar">
                    <div class="h-8 w-1.5 rounded-full bg-yellow-400"></div>
                    <div class="flex w-full justify-between text-sm font-medium">
                        <span>Estimated Wait Time:</span>
                        <span>{estimated_wait_time_minutes} Minutes</span>
                    </div>
                </div>
            </div>
    
            <Drawer.Header>
                <Drawer.Title class="text-2xl">Make a Live Report</Drawer.Title>
                <Drawer.Description class="text-sm"
                    >Help other players see if the courts are busy</Drawer.Description
                >
            </Drawer.Header>
    
            <div class="flex flex-col gap-2 p-2 px-2 pb-0">
                <!-- Courts Occupied -->
                <label for="courts" class="text-sm font-medium">
                    Courts Occupied
                </label>
                <input
                    id="courts"
                    type="number"
                    min="0"
                    class="bg-background focus:ring-primary w-full rounded-md border border-neutral-300 px-2 py-1 text-sm focus:outline-none focus:ring-2"
                    placeholder="e.g. 2"
                />
    
                <!-- Groups Waiting -->
                <label for="groups" class="text-sm font-medium"> Groups Waiting </label>
                <input
                    id="groups"
                    type="number"
                    min="0"
                    class="bg-background focus:ring-primary w-full rounded-md border border-neutral-300 px-2 py-1 text-sm focus:outline-none focus:ring-2"
                    placeholder="e.g. 1"
                />
                <button class="self-start bg-neutral-200 p-2 rounded-xl">
                    Submit
                </button>
            </div>
        </div>
	</Drawer.Content>
</Drawer.Root>

<style lang="postcss">
	@reference "tailwindcss";

	.bar {
		@apply border-1 flex items-center gap-3 rounded-lg border-neutral-300 p-2 pr-4;
	}
</style>
