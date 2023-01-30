<?php

namespace App\Controllers;

use Spatie\Dns\Dns;

class Home extends BaseController
{
	private $ip, $reverse;

	private function getIpInfo() {
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
		} catch (Exception $e) {
			$reverse = "";
		}

		if (empty($reverse)) {
			$reverse = "";
		} else {
			$reverse = end($reverse);
			$reverse = $reverse->target();
		}

		if ($reverse && str_ends_with($reverse, '.')) {
			$reverse = substr($reverse, 0, -1);
		}

		$this->ip = $ip;
		$this->reverse = $reverse;
	}

	public function index()
	{
		$this->getIpInfo();

		try {
			$bing = new \grubersjoe\BingPhoto();
			$background = $bing->getImage();
		} catch(Exception $e) {
			$background = null;
		}

	
		return $this->twig->display('home', [
			'ip' => $this->ip,
			'reverse' => $this->reverse,
			'background' => $background
		]);
	}

	public function json()
	{
		$this->getIpInfo();

		return $this->response->setJSON([
			'ip' => $this->ip,
			'reverse' => $this->reverse
		]);
	}

	public function xml()
	{
		$this->getIpInfo();

		return $this->response->setXML([
			'ip' => $this->ip,
			'reverse' => $this->reverse
		]);
	}

	public function ip()
	{
		$this->getIpInfo();

		$this->response->setContentType('text/plain');
		echo "{$this->ip}\n";
	}

	public function reverse()
	{
		$this->getIpInfo();

		$this->response->setContentType('text/plain');
		echo "{$this->reverse}\n";
	}
}
