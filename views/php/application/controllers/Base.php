<?php 

class BaseController extends Yaf_Controller_Abstract {
    public function init () {
        Yaf_Dispatcher::getInstance()->disableView(); 
    }

    public function __errResponseJson($errcode = -1, $errmsg = 'failed') {
        $res = array('errcode' => $errcode, 'errmsg' => $errmsg);
        $this->getResponse()->setHeader('Content-Type','json;charset=UTF-8');
        $this->getResponse()->setHeader('Server','Zat');
        $this->getResponse()->setHeader('X-Powered-By','Zat');
        $this->getResponse()->setBody(json_encode($res));
    }

    public function __responseJson($data = NULL, $errcode = 0, $errmsg = 'success') {
        $res = array('errcode' => $errcode, 'errmsg' => $errmsg, 'data' => $data);
        $this->getResponse()->setHeader('Content-Type','json;charset=UTF-8');
        $this->getResponse()->setHeader('Server','Zat');
        $this->getResponse()->setHeader('X-Powered-By','Zat');
        $this->getResponse()->setBody(json_encode($res));
    }
}

?>