// @ts-nocheck
import { Route } from "wouter";
import SimpleLayout from "../layouts/shell/simple";
import {{title project}}Layout from "../layouts/shell/{{toLower project}}";

// SimpleRoute wraps the route with a simple parent layout.
export const SimpleRoute = ({ path, header, content }) => {
	return (
		<Route
			path={path}
			component={(props) => (
				<SimpleLayout header={header} content={content} {...props} />
			)}
		/>
	);
};

// ComplexRoute wraps the route with a complex parent layout.
export const ComplexRoute = ({ path, header, content }) => {
	return (
		<Route
			path={path}
			component={(props) => (
				<{{title project}}Layout header={header} content={content} {...props} />
			)}
		/>
	);
};
