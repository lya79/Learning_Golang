//
//  ViewController.swift
//  HelloWorldCallback
//
//  Created by yajun on 18/07/2018.
//  Copyright © 2018 ising99. All rights reserved.
//

import UIKit
import Hello

class ViewController: UIViewController, HelloEventListenerProtocol{
    
    
    var output = HelloOutput() {
        didSet {
            print(output?.msg())
        }
    }
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    @IBAction func showMessage(sender: UIButton){
        let input = Hello.HelloInput()
        input?.setName("Hello World")
        
        HelloHello(input, self)
    }
    
    func callback(_ output: HelloOutput!) {
        var msg = output.msg()
        let alertController = UIAlertController(title:"Welcome to my first app", message:msg, preferredStyle: UIAlertControllerStyle.alert)
        alertController.addAction(UIAlertAction(title:"OK", style: UIAlertActionStyle.default, handler:nil))
        present(alertController, animated:true, completion: nil)
    }
}

