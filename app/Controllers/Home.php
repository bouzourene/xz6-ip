<?php

namespace App\Controllers;

use Spatie\Dns\Dns;

class Home extends BaseController
{
	public function index()
	{
		$ip = $this->request->getIPAddress();

		if (str_contains($ip, '.')) {
			$arpa = explode('.', $ip);
			$arpa = array_reverse($arpa);
			$arpa = implode('.', $arpa);
			$arpa = "{$arpa}.in-addr.arpa";
		} else {
			$address = \IPLib\Address\IPv6::parseString($ip);
			$address = str_replace(':', '', $address->toString(true));

			$arpa = str_split($address);
			$arpa = array_reverse($arpa);
			$arpa = implode('.', $arpa);
			$arpa = "{$arpa}.ip6.arpa";
		}

		$dns = new Dns();
		try {
			$reverse = $dns->getRecords($arpa, 'PTR');
			$reverse = end($reverse);
			$reverse = $reverse->target();
		} catch (Exception $e) {
			$reverse = null;
		}

		if ($reverse && str_ends_with($reverse, '.')) {
			$reverse = substr($reverse, 0, -1);
		}

		$bing = new \grubersjoe\BingPhoto();
		$background = $bing->getImage();
	
		return $this->twig->display('home', [
			'ip' => $ip,
			'reverse' => $reverse,
			'background' => $background
		]);
	}
}
