import { h, Component } from 'preact';
import { connect } from 'preact-redux';
import reduce from '../reducers';
import * as actions from '../actions';
// import Users from './users';
// import TodoItem from './todo-item';

import UserServices from '../services/user.servies'

import LayoutGrid from 'preact-material-components/LayoutGrid';
import 'preact-material-components/LayoutGrid/style.css';
import Dialog from 'preact-material-components/Dialog';
import Button from 'preact-material-components/Button';
import List from 'preact-material-components/List';
import 'preact-material-components/List/style.css';
import 'preact-material-components/Button/style.css';
import 'preact-material-components/Dialog/style.css';
import Radio from 'preact-material-components/Radio';
import FormField from 'preact-material-components/FormField';
import 'preact-material-components/Radio/style.css';
import 'preact-material-components/FormField/style.css';
import TextField from 'preact-material-components/TextField';
import 'preact-material-components/TextField/style.css';
import Snackbar from 'preact-material-components/Snackbar';
import 'preact-material-components/Button/style.css';
import 'preact-material-components/Snackbar/style.css';
import Card from 'preact-material-components/Card';
import 'preact-material-components/Card/style.css';
import 'preact-material-components/Button/style.css';
import Chips from 'preact-material-components/Chips';
import 'preact-material-components/Chips/style.css';
import 'preact-material-components/Theme/style.css';
import TopAppBar from 'preact-material-components/TopAppBar';
import 'preact-material-components/TopAppBar/style.css';

@connect(reduce, actions)

export default class App extends Component {

	emptyUserObject = {
		firstName: "",
		lastName: "",
		skills: "",
		interests: "",
	};
	
	state = {
		dialogHeaderText: "Add User",
		dialogUpdateText: "Create",
		deleteDialogBody: "",
		deleteDialogHeader: "",
		dialogErrorText: "",
		selectedUser: "",
		user: {},
		users: [],
	};

	styles = {
		input : {
			"padding-left": "18px",
			"paddign-top": "10px"
		}
	}

	updateUser = async () => {
		let msg = this.validateForm();
		if (msg) {
			this.showMessage(msg);
		} else {
			console.log(this.state.selectedUser)
			if (!this.state.selectedUser) {
				let user = await UserServices.addUser(this.state.user);
				console.log(user);
				this.props.addUser(user)
			} else {
				let user = await UserServices.updateUser(this.state.user);
				console.log(user);
				this.props.updateUser(user)
			}
			this.hideUserDialog();
		}
	}

	validateForm = () => {
		let msg = "";
		if (!this.state.user.firstName) {
			msg = "Please provide valid first name";
		} else if (!this.state.user.lastName) {
			msg = "Please provide valid last name";
		} else if (!this.state.user.skills) {
			msg = "Please provide atleast one skill";
		} else if (!this.state.user.interests) {
			msg = "Please provide atleast one interest";
		}
		return msg;
	}

	handleChange = (e, key) => {
		this.setState({
			user: {
				...this.state.user,
				[key]: e.target.value
			}
		});
	}

	showMessage = (msg) => {
		this.bar.MDComponent.show({
            message: msg
		});
	}

	openUpdateUserDialog = user => {
		if (!user) {
			this.setState({
				selectedUser: "",
				dialogHeaderText: "Add User",
				dialogUpdateText: "Create",
				user : Object.assign({}, this.emptyUserObject),
			});
		} else {
			this.setState({
				selectedUser: user.id,
				dialogUpdateText: "Update",
				dialogHeaderText: `Edit ${user.firstName}`,
				user : Object.assign({}, user),
			});
		}
		this.showUserDialog();
	};

	openDeleteUserDialog = user => {
		this.setState({
			selectedUser: user.id,
			deleteDialogHeader: "Delete",
			deleteDialogBody: `Are you sure you want to delete user - ${user.firstName} ${user.lastName}`,
		});
		this.delDialog.MDComponent.show()
	}

	deleteUser = async () => {
		if (this.state.selectedUser) {
			let user = await UserServices.deleteUser(this.state.selectedUser);
			this.props.removeUser(this.state.selectedUser)
			this.setState({
				selectedUser: "",
			});
		}
		this.delDialog.MDComponent.close()
	}

	showUserDialog = () => {
		this.dialog.MDComponent.show()
	}

	hideUserDialog = () => {
		this.dialog.MDComponent.close()
	}

	componentDidMount = async () => {
		let users = await UserServices.getAllUser();
		users.forEach(this.props.addUser);
	}

	render({ users }) {
		return (
			<div id="app">
				<TopAppBar className="topappbar">
					<TopAppBar.Row>
					<TopAppBar.Section align-start>
						<TopAppBar.Title>
						Kube Safe - the Users
						</TopAppBar.Title>
					</TopAppBar.Section>
					<TopAppBar.Section align-end>
						<Button id="addUserButton" secondary={true} raised={true} onClick={() => this.openUpdateUserDialog()}>
						<span>Add New User</span>
						</Button>
					</TopAppBar.Section>
					</TopAppBar.Row>
				</TopAppBar>
				<LayoutGrid>
					<LayoutGrid.Inner>
						<LayoutGrid.Cell cols="12" align="middle">
						
							<Dialog ref={dialog=>{this.dialog=dialog;}}>
								<Dialog.Header>{this.state.dialogHeaderText}</Dialog.Header>
								<Dialog.Body scrollable={true}>
									<FormField style={{"width":"100%"}}>
										<TextField label="First Name" fullwidth={true} box={true} value={this.state.user.firstName} onChange={(e) => this.handleChange(e, "firstName")}/>
										<TextField label="Last Name" fullwidth={true} value={this.state.user.lastName} onChange={(e) => this.handleChange(e, "lastName")}/>
									</FormField>
									<FormField style={{"width":"100%"}}>
										<TextField style={{"width":"100%"}} label="Skill(s) comma separated" fullwidth={true} box={true} value={this.state.user.skills} onChange={(e) => this.handleChange(e, "skills")}/>
									</FormField>
									<FormField style={{"width":"100%"}}>
										<TextField style={{"width":"100%"}} label="Interests(s) comma separated" fullwidth={true} box={true} value={this.state.user.interests} onChange={(e) => this.handleChange(e, "interests")}/>
									</FormField>
								</Dialog.Body>
								<Dialog.Footer>
									<Dialog.FooterButton cancel={true}>Cancel</Dialog.FooterButton>
									<Dialog.FooterButton onClick={() => this.updateUser()}>{this.state.dialogUpdateText}</Dialog.FooterButton>
								</Dialog.Footer>
							</Dialog>
						</LayoutGrid.Cell>
						
					</LayoutGrid.Inner>
				</LayoutGrid>
				

				<LayoutGrid>
					<LayoutGrid.Inner>
						<LayoutGrid.Cell cols="12" align="middle" style={{"text-align": "center"}}>
						{
							users.length ? users.map(user => (
								// <div>{user.id}</div>
								<Card>
									<div class="card-header">
										<h2 class=" mdc-typography--title">{user.firstName} {user.lastName}</h2>
									</div>
									<Card.Media className="card-media" style="text-align:left;">
									<h4 class=" mdc-typography--title">
										<span style="margin-left: 10px">Skills: </span>
										<Chips>
											{
												user.skills.split(',').map(s => {
													if (s.trim().length != 0) {
														return (
															<Chips.Chip>
																<Chips.Text>{s}</Chips.Text>
															</Chips.Chip>
														)
													} else {
														return <div></div>
													}
												})
											}
										</Chips>
										</h4>

										<h4 class=" mdc-typography--title">
										<span style="margin-left: 10px">Interests: </span>
										<Chips>
											{
												user.interests.split(',').map(s => {
													if (s.trim().length != 0) {
														return (
															<Chips.Chip>
																<Chips.Text>{s}</Chips.Text>
															</Chips.Chip>
														)
													} else {
														return <div></div>
													}
												})
											}
										</Chips>
										</h4>
									</Card.Media>
									<Card.Actions>
										<Card.ActionButtons>
											<Card.ActionButton onClick={() => this.openUpdateUserDialog(user)}>Edit</Card.ActionButton>
											<Card.ActionButton onClick={() => this.openDeleteUserDialog(user)}>Delete</Card.ActionButton>
										</Card.ActionButtons>
									</Card.Actions>
								</Card>
							)) : <h2>No Rows</h2>
						}
						</LayoutGrid.Cell>
					</LayoutGrid.Inner>
				</LayoutGrid>
				<Dialog ref={delDialog=>{this.delDialog=delDialog;}}>
					<Dialog.Header>{this.state.deleteDialogHeader}</Dialog.Header>
					<Dialog.Body>
						{
							this.state.deleteDialogBody
						}
					</Dialog.Body>
					<Dialog.Footer>
						<Dialog.FooterButton cancel={true}>Decline</Dialog.FooterButton>
						<Dialog.FooterButton onClick={() => this.deleteUser()}>Delete</Dialog.FooterButton>
					</Dialog.Footer>
				</Dialog>
				<Snackbar ref={bar=>{this.bar=bar;}}/>
			</div>
		);
	}
}
