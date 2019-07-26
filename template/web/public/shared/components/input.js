import styles from "./input.module.css";
import classnames from "classnames";

export default (props) => (
	<input
		type={props.type || "text"}
		ref={props.ref}
		value={props.value}
		className={classnames(styles.root, props.className)}
		onClick={props.onClick}
		onChange={props.onChange}
		onMouseEnter={props.onMouseEnter}
		onMouseLeave={props.onMouseLeave}
		disabled={props.disabled}
		spellcheck={props.spellCheck}
		placeholder={props.placeholder}
	/>
);
