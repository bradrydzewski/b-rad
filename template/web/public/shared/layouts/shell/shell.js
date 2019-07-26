import styles from "./shell.module.css";
import classnames from "classnames";

import Link from "../../link.js";

export default (props) => (
	<>
		<nav>
			<ul>
				<li>
					<Link href="/">Home</Link>
				</li>
				{props.session.user.admin ? (
					<li>
						<Link href="/users">Users</Link>
					</li>
				) : undefined}
				<li>
					<Link href="/account">Account</Link>
				</li>
			</ul>
		</nav>
		<div className={classnames(styles.root, props.className)}>
			{props.children}
		</div>
	</>
);
