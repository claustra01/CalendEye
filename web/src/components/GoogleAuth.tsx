import liff from '@line/liff';

type Props = {
	userId: string;
	displayName: string;
};

function GoogleLogin({ props }: { props: Props }) {
	const redirectToWeb = () => {
		liff.openWindow({
			url: `/oauth2?id=${props.userId}`,
			external: true,
		});
	};

	return (
		<>
			<p>ようこそ、{props.displayName}</p>
			<button type="submit" onClick={() => redirectToWeb()}>
				Login with Google
			</button>
		</>
	);
}

export default GoogleLogin;
