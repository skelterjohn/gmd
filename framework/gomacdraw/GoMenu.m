//
//  GoMenu.m
//  gomacdraw
//
//  Created by John Asmuth on 5/17/11.
//  Copyright 2011 Rutgers University. All rights reserved.
//

#import "GoMenu.h"


@implementation GoMenu

@synthesize menu;

- (id)init
{
    self = [super init];
    if (self) {
        // Initialization code here.
    }
    
    return self;
}

- (void)dealloc
{
    [super dealloc];
}

- (void)setAppName:(NSString*)name
{
    NSString* oldTitle = @"Quit go";
    NSString* newTitle = [NSString stringWithFormat:@"Quit %@", name];
    [[[[menu itemAtIndex:0] submenu] itemWithTitle:oldTitle] setTitle:newTitle];
}

- (void)loadNoNib
{
    menu = [NSApp mainMenu];
    if (menu == NULL) {
        NSLog(@"menu is null");
    }
    if (NSApp == NULL) {
        NSLog(@"NSApp is null");
    }
    
    NSMenuItem* appItem = [[NSMenuItem allocWithZone:[NSMenu menuZone]] initWithTitle:@"Go" action:NULL keyEquivalent:@""];
    
    NSMenu* appMenu = [[NSMenu alloc] initWithTitle:@"Go"];
    [appMenu insertItemWithTitle:@"About go" action:NULL keyEquivalent:@"" atIndex:0];
    [appMenu insertItemWithTitle:@"Quit go" action:NULL keyEquivalent:@"" atIndex:1];
    
    [appItem setSubmenu:appMenu];
    
    [menu addItem:appItem];
    
    NSMenuItem* fileItem = [[NSMenuItem allocWithZone:[NSMenu menuZone]] initWithTitle:@"File" action:NULL keyEquivalent:@""];
    [menu addItem:fileItem];
}

@end
