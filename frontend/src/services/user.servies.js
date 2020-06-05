/* eslint-disable no-console */
import axios from 'axios';

const API = `http://localhost:2369/api`;

const UserServices = {
	addUser: async user => {
		try {
			let res = await axios.post(`${API}/createUser`, user);
			return res.data ? res.data.result : null;
		}
		catch (err) {
			console.log(err);
			return null;
		}
	},
	updateUser: async user => {
		try {
			let res = await axios.put(`${API}/updateUser/${user.id}`, user);
			return res.data ? res.data.result : null;
		}
		catch (err) {
			console.log(err);
			return null;
		}
	},
	deleteUser: async id => {
		try {
			let res = await axios.delete(`${API}/deleteUser/${id}`);
			return res.data ? res.data.result : null;
		}
		catch (err) {
			console.log(err);
			return null;
		}
	},
	getAllUser: async () => {
		try {
			let res = await axios.get(`${API}/getAllUsers`);
			return res.data ? res.data.result : null;
		}
		catch (err) {
			console.log(err);
			return null;
		}
	}
};

export default UserServices;