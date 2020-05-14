#include <iostream>

#include "wavmatch.h"
#include <stdlib.h>
#include <mutex>
#include <memory>
#include "yispeech.h"
#include <string.h>
#include <dirent.h>
#include <sys/stat.h>

#define MAX_LEN (255)

static std::mutex g_locker;

std::vector<std::string> g_templates;

void loadTemplates(const char * path, std::vector<std::string>& templates);

//load all template into memory
//path: template path
int ReloadTemplate(const char* path){
    size_t len = strlen(path);
    if ((0 > len) || (MAX_LEN < len))
    {
        //无效的path
        return -1;
    }

    int numbers;
    std::vector<std::string> templates;
    loadTemplates(path, templates);
    if (0 == templates.size())
    {
        //loadTemplate 失败
        std::cout << "templates load failed, from: " << path <<std::endl;
        return -2;
    }

    {
        std::unique_lock<std::mutex> lock(g_locker);
        g_templates= templates;
    }

    return 0;
}

//match a wav to get end reason
//wav_path: wav path
//matched_path: matched template path
int WavMatchCos(const char* cstrWavPath, char** cstrMatchedPath){
    std::string wavPath(cstrWavPath);

    std::vector<std::string> vtemplates;
    {
        std::unique_lock<std::mutex> lock(g_locker);
        vtemplates = g_templates;
    }

    //do match
    std::vector<float> cos = yispeech::WavMatchCos(wavPath, vtemplates);
    if (cos.size() != vtemplates.size()){
        return -2;
    }

    //get biggest from cos ->matchIndex
    size_t matchIndex = 0;
    float max = 0.0;
    size_t i=0;
    for (i=0; i< vtemplates.size(); i++) {
        if (cos[i] > max) {
            matchIndex = i;
            max = cos[i];
        }
    }

    if (max < 0.90) {
        std::cout << "max:" << max << std::endl;
        return vtemplates.size();
    }

    //get cstrMatchedPath by matchIndex and vtemplates;
    char* matchpath = (char *)malloc(vtemplates[matchIndex].length());
    memcpy((void*)matchpath, vtemplates[matchIndex].c_str(), vtemplates[matchIndex].length());
    *cstrMatchedPath = matchpath;

    return 0;
}

 void loadTemplates(const char * path, std::vector<std::string>& templates) {
    DIR* dirp = opendir(path);
    if(!dirp)
    {
        return ;
    }

    struct stat st;
    struct dirent *dir;
    while((dir = readdir(dirp)) != NULL)
    {
        if(strcmp(dir->d_name,".") == 0 ||
                   strcmp(dir->d_name,"..") == 0)    
        {
            continue;
        }

        std::string name = dir->d_name;
        if (std::string::npos == name.find(".wav")) {
            continue;
        }
        
        std::string full_path = path + name;
        if(lstat(full_path.c_str(),&st) == -1)
        {
            continue;
        }

        if(!S_ISDIR(st.st_mode))   //S_ISDIR()宏判断是否是目录文件
        {
            templates.push_back(full_path);
        }
    }
    closedir(dirp);

}
