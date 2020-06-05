import { createStore } from 'redux';

let ACTIONS = {
	ADD_USER: ({ users, ...state }, { user }) => ({
		users: [user, ...users],
		...state
	}),
	UPDATE_USER: ({ users, ...state }, { user }) => {
		let allUsers = users.map(u => {
			if (u.id === user.id) {
				return user;
			}
			return u;
		});
		return {
			users: allUsers,
			...state
		};
	},
	REMOVE_USER: ({ users, ...state }, { id }) => ({
		users: users.filter( i => i.id!==id ),
		...state
	})
};

const INITIAL = {
	users: []
};

export default createStore( (state, action) => (
	action && ACTIONS[action.type] ? ACTIONS[action.type](state, action) : state
), INITIAL, typeof devToolsExtension==='function' ? devToolsExtension() : undefined);
