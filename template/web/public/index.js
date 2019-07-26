import { render } from "react-dom";
import { Route, Switch, Redirect } from "wouter";
import { SWRConfig } from "swr";

import { ProvideSession, useSession } from "./hooks/session.js";

import Account from "./pages/account.js";
import Home from "./pages/home.js";
import Login from "./pages/login.js";
import Project from "./pages/project.js";
import Register from "./pages/register.js";
import Users from "./pages/users.js";

import Shell from "./shared/layouts/shell/shell.js";
import Guest from "./shared/layouts/login.js";

// TODO remove me
import Demo from "./shared/components/demo/demo.js";

export default function App() {
	const { session, fetcher } = useSession();

	// if the session is loaded, and the session
	// is falsey, the login and register routes
	// are exposed.
	if (!session) {
		return (
			<>
				<Guest>
					<Switch>
						<Route path="/demo" component={Demo} />
						<Route path="/register" component={Register} />
						<Route component={Login} />
					</Switch>
				</Guest>
			</>
		);
	}

	return (
		<>
			<Shell session={session}>
				<SWRConfig value={{`{{`}} fetcher {{`}}`}}>
					<Switch>
						<Route path="/" component={Home} />
						<Route path="/users" component={Users} />
						<Route path="/projects/:project" component={Project} />
						<Route path="/projects/:project/:path+" component={Project} />
						<Route path="/account" component={Account} />
						<Route path="/login">
							<Redirect to={"/"} />
						</Route>
						<Route path="/register">
							<Redirect to={"/"} />
						</Route>
						<Route>Not Found</Route>
					</Switch>
				</SWRConfig>
			</Shell>
		</>
	);
}

render(
	<ProvideSession>
		<App />
	</ProvideSession>,
	document.body
);
