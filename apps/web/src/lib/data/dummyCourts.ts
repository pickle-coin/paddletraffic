// Dummy court data matching the OpenAPI Court schema
// This will be replaced with real API calls later

export interface Court {
	id: number;
	name: string;
	courtCount: number;
	location: {
		addressLine: string;
		region?: string;
		postalCode?: string;
		countryCode: string;
		timezone: string;
		coordinates: {
			lat: number;
			lon: number;
		};
		placeId?: string;
	};
	status: {
		courtsOccupied: number;
		groupsWaiting: number;
		lastReport: string;
		courtId: number;
	};
}

export const dummyCourts: Court[] = [
	{
		id: 1,
		name: '11th Ave Park',
		courtCount: 6,
		location: {
			addressLine: '581 Terrace Hills Dr',
			region: 'UT',
			postalCode: '84103',
			countryCode: 'US',
			timezone: 'America/Denver',
			coordinates: {
				lat: 40.783488,
				lon: -111.862136
			}
		},
		status: {
			courtsOccupied: 0,
			groupsWaiting: 0,
			lastReport: '2025-10-26T10:00:00Z',
			courtId: 1
		}
	},
	{
		id: 2,
		name: 'Fairmont Park',
		courtCount: 6,
		location: {
			addressLine: '2305 900 E',
			region: 'UT',
			postalCode: '84106',
			countryCode: 'US',
			timezone: 'America/Denver',
			coordinates: {
				lat: 40.720150,
				lon: -111.862812
			}
		},
		status: {
			courtsOccupied: 0,
			groupsWaiting: 0,
			lastReport: '2025-10-26T10:00:00Z',
			courtId: 2
		}
	},
	{
		id: 3,
		name: '5th Ave & C Street Pickleball Courts',
		courtCount: 2,
		location: {
			addressLine: '230 C St E',
			region: 'UT',
			postalCode: '84103',
			countryCode: 'US',
			timezone: 'America/Denver',
			coordinates: {
				lat: 40.774842,
				lon: -111.880207
			}
		},
		status: {
			courtsOccupied: 0,
			groupsWaiting: 0,
			lastReport: '2025-10-26T10:00:00Z',
			courtId: 3
		}
	},
	{
		id: 4,
		name: 'Rosewood Park',
		courtCount: 8,
		location: {
			addressLine: '1400 N 1200 W',
			region: 'UT',
			countryCode: 'US',
			timezone: 'America/Denver',
			coordinates: {
				lat: 40.800008,
				lon: -111.925006
			}
		},
		status: {
			courtsOccupied: 0,
			groupsWaiting: 0,
			lastReport: '2025-10-26T10:00:00Z',
			courtId: 4
		}
	},
	{
		id: 5,
		name: 'Canyon Rim Park Pickleball Courts',
		courtCount: 10,
		location: {
			addressLine: '3100 S Grace St',
			region: 'UT',
			postalCode: '84109',
			countryCode: 'US',
			timezone: 'America/Denver',
			coordinates: {
				lat: 40.705049,
				lon: -111.808643
			}
		},
		status: {
			courtsOccupied: 0,
			groupsWaiting: 0,
			lastReport: '2025-10-26T10:00:00Z',
			courtId: 5
		}
	}
];
