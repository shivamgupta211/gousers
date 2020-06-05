
export function addUser(user) {
	return {
		type: 'ADD_USER',
		user
	};
}

export function updateUser(user) {
	return {
		type: 'UPDATE_USER',
		user
	};
}

export function removeUser(id) {
	return {
		type: 'REMOVE_USER',
		id
	};
}
