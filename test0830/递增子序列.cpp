#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include <assert.h>
#include <queue>

using namespace std;

bool IsString(char *str, int len)
{
    if(str == NULL || len < 3)
    {
        return false;
    }
    int i = 0;
    int num = 0;
    
    for(i=0; i<len-1; i++)
    {
        
        if(str[i]<str[i+1])
        {
            num++;
        }
        else if(str[i]>=str[i+1]) 
        {
            num = 0;
        }
        else if(num == 3)
        {
            return true;
        }
    }
    return false;
}

int main()
{
    int n;
    cin>>n;
    char *str = (char*)malloc(sizeof(char)*n);

    if(n < 3)
    {
        return 0;
    }
    for(int i=0; i<n; ++i)
    {
        cin>>str[i];
    }
    bool s = IsString(str, n);
    if(s == 0)
    {
        cout<<false<<endl;
    }
    else
    {
        cout<<true<<endl;
    }
    
    
    return 0;
}