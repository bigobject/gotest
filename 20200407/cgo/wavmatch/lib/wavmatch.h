#ifndef _WAV_MATCH_H
#define _WAV_MATCH_H
#ifdef __cplusplus
extern "C" {
#endif

//load all template into memory
//path: template path
//return 0 if success, otherwise return error code
int ReloadTemplate(const char* path);

//match a wav to get end reason
//wav_path: wav path
//matched_path: matched template path
//return 0 if success, otherwise return error code
int WavMatchCos(const char* wav_path, char** matched_path);
 
#ifdef __cplusplus    
}
#endif

#endif
 
