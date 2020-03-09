<?php

	require_once('WebToPay.php');

	function get_self_url() {
		$s = substr(strtolower($_SERVER['SERVER_PROTOCOL']), 0,
		            strpos($_SERVER['SERVER_PROTOCOL'], '/'));

		if (!empty($_SERVER["HTTPS"])) {
		    $s .= ($_SERVER["HTTPS"] == "on") ? "s" : "";
		}

		$s .= '://'.$_SERVER['HTTP_HOST'];

		if (!empty($_SERVER['SERVER_PORT']) && $_SERVER['SERVER_PORT'] != '80') {
		    $s .= ':'.$_SERVER['SERVER_PORT'];
		}

		$s .= dirname($_SERVER['SCRIPT_NAME']);

		return $s;
	}

	try {
		$self_url = get_self_url();

		$request = WebToPay::buildRequest(array(
		        // Čia surašyti tik keli parametrai.
		        // Visų galimų parametrų su aprašymais sąrašą rasite žemiau.
		        'projectid'     => 81824,
		        'sign_password' => '856b1cacaea25f6ebbf01b0003030a0b',
				'orderid'		=> 0,
		        'country'       => 'IE',
            	'lang'          => 'ENG',
				'paytext'       => 'Dotacja.',
		        'accepturl'     => $self_url.'http://nforum.pl/libwebtopay/accept.php',
		        'cancelurl'     => $self_url.'http://nforum.pl/libwebtopay/cancel.php',
		        'callbackurl'   => $self_url.'/callback.php',
		        'test'          => 0,
		    ));
	} catch (WebToPayException $e) {
		echo $e->getMessage();
	}

?>
<form method="post" action="<?php echo WebToPay::PAY_URL; ?>">
    <?php foreach ($request as $key => $val): ?>
    <input type="hidden" name="<?php echo $key ?>"
           value="<?php echo get_magic_quotes_gpc() ? $val : addslashes($val); ?>" /> 
    <?php endforeach; ?>
    <input type="image" border="0" name="submit"
           src="https://static1.squarespace.com/static/507de88a84ae4cfb9e22879b/507de94184ae4cfb9e228a35/50c13079e4b0aab37deccaf7/1354838138959/Smaller+Donate+Button@2x.png" width="530"
           alt="Paremkite svetainę! - mokejimai.lt tai saugu ir patikima!" />
</form>
