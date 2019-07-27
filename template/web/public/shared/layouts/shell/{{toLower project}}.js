import styles from "./{{toLower project}}.module.css";
import classnames from "classnames";
import { useSession } from "../../../hooks/session.js";

import { useLocation } from "wouter";

import Link from "../../link.js";
import Avatar from "../../components/avatar";
import Button from "../../components/button";

export default ({ sidebar: Sidebar, content: Content, className, ...rest }) => {
	const { session } = useSession();
	const [location, setLocation] = useLocation();

	return (
		<main className={styles.root}>
			<aside className={styles.sidebar}>
				<div className={styles.nav}>
					<ul>
						<li>
							<Link href={`/{{toLower project}}s/${rest.params.{{toLower project}}{{`}`}}`}>{{title parent}}s</Link>
						</li>
						<li>
							<Link href={`/{{toLower project}}s/${rest.params.{{toLower project}}{{`}`}}/members`}>
								Members
							</Link>
						</li>
						<li>
							<Link href={`/{{toLower project}}s/${rest.params.{{toLower project}}{{`}`}}/settings`}>
								Settings
							</Link>
						</li>
					</ul>
				</div>
			</aside>
			<header className={styles.header}>
				<Link href="/">
					<img src="/logo.svg" />
				</Link>
				<div className={styles.menu}>
					{session.user.admin ? (
						<div className={styles.logout}>
							<Button onClick={() => setLocation("/users")}>Admin</Button>
						</div>
					) : undefined}
				</div>
				<Link href="/account">
					<Avatar text={session.user.email} />
				</Link>
			</header>

			<div className={classnames(styles.content, className)}>
				<Content {...rest} />
			</div>
		</main>
	);
};
