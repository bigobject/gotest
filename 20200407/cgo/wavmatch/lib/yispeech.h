//
// Created by rukuang on 2018/7/3.
//

#ifndef YISPEECH_YISPEECH_H_
#define YISPEECH_YISPEECH_H_

#include <memory>
#include <string>
#include <tuple>
#include <utility>
#include <vector>

namespace yispeech {
/**
 * @brief  资源共享类，用于在线识别器
 * @note
 * 使用在线识别器前，需要先声明此类，并且保持此类的生命周期与在线识别器一样
 */
class SpeechResource {
 public:
  /**
   * @brief  初始化
   * @note   与在线识别器的配置文件相同
   * @param  config_path: 配置文件的路径
   */
  explicit SpeechResource(const std::string &config_path);

  friend class OnlineSpeechRecognizer;

  ~SpeechResource();

 private:
  class ResourceImpl;  // !< 具体实现类

  std::shared_ptr<ResourceImpl> p_impl_;  // !< 封装了具体实现类的指针
};

/**
 * @brief 非在线的语音识别器.
 * @note 此对象构建开销较大，在识别过程中需要一直保持
 */
class SpeechRecognizer {
 public:
  /**
   * @brief 指定语音识别器的相关配置, 初始化识别器.
   * @param config_path yaml格式的配置文件.
   */
  explicit SpeechRecognizer(const std::string &config_path);

  /**
   * @brief 识别语音文件.
   * @param wav_path 需要识别的音频文件, 必须是wav格式.
   * @return 识别的字符串.
   */
  std::string Recognize(const std::string &wav_path) const;

  ~SpeechRecognizer();

 private:
  class SpeechRecognizerImpl;  // !< 具体实现类

  std::unique_ptr<SpeechRecognizerImpl> p_impl_;  // !< 封装了具体实现类的指针
};

/**
 * @brief 非在线的并行语音识别器.
 * @note 此对象构建开销较大，在识别过程中需要一直保持
 */
class ParallelSpeechRecognizer {
 public:
  /**
   * @brief 指定语音识别器的相关配置, 初始化识别器.
   * @param config_path yaml格式的配置文件.
   */
  explicit ParallelSpeechRecognizer(const std::string &config_path);

  /**
   * @brief 并行识别语音文件.
   * @param wav_paths 需要识别的音频文件, 必须是wav地址组成的vector.
   * @return 识别的字符串vector，顺序与地址的顺序对应.
   */
  std::vector<std::string> Recognize(
      const std::vector<std::string> &wav_paths) const;

  /**
   * @brief 结合说话人信息识别语音文件.
   * @param spkr_wav_pairs 与说话人相关的音频文件容器，维度是[说话人，一句话的地址]
   * @return 与wav_paths结构相同的识别结果，维度是[说话人，一句话的识别结果]
   */
  std::vector<std::vector<std::string>> RecognizeWithSpeaker(
      const std::vector<std::vector<std::string>> &spkr_wav_pairs) const;

  ~ParallelSpeechRecognizer();

 private:
  class PlSpeechRecognizerImpl;  // !< 具体实现类

  std::unique_ptr<PlSpeechRecognizerImpl> p_impl_;  // !< 封装了具体实现类的指针
};

/**
 * @brief 在线语音识别器.
 * @note 此对象构建开销较大，在识别过程中需要一直保持
 */
class OnlineSpeechRecognizer {
 public:
  /**
   * @brief 目前只允许资源不共享的形式提供在线语音.
   * 如果多线程测试通过之后可能改变接口的形态.
   * @param conf_path yaml格式的配置文件路径.
   * @param speech_res 资源共享类
   */
  explicit OnlineSpeechRecognizer(const std::string &conf_path,
                                  const SpeechResource &speech_res);

  /**
   * @brief 对一个片段进行识别, 目前该值的返回和类的状态相关.
   * @note 仅支持8k频率，PCM位宽为2的数据块。
   * @param wav_data wav音频的数据段, 目前硬性要求180ms长度.
   * @param is_end 该音频段是否是最后一个音频段.
   * @return 整个音频到目前为止的识别结果.
   */
  std::string Recognize(const std::vector<char> &wav_data, bool is_end) const;

  /**
   * @brief 调用此方法清空对象内状态
   * @note 此方法以最小的代价最快使对象清空，但放弃了之前输入的音频流的识别结果
   */
  void Reset() const;

  ~OnlineSpeechRecognizer();

 private:
  class OlSpeechRecognizerImpl;  // !< 具体实现类
  std::unique_ptr<OlSpeechRecognizerImpl> ol_p_impl_;
};

/**
 * @brief 在线Vad类.
 * @note 仅支持8k采样频率, 16bit的音频格式.
 */
class OnlineVad {
 public:
  /**
   * @brief 构造.
   * @param vad_yaml vad配置文件路径, 参考vad.yaml.
   */
  explicit OnlineVad(const std::string &vad_yaml);

  /**
   * 移动构造函数
   * @param rhs 需要移动的类.
   */
  OnlineVad(OnlineVad &&rhs) noexcept;

  /**
   * @brief 重置, 清空所有缓存.
   */
  void Reset() const;

  /**
   * @brief 检测Vad端点.
   * @param wav_data wav数据, 建议200ms长度, 仅支持8k频率, 16bit的数据块.
   * @param is_end 该音频段是否是为最后一片.
   * @return 时间端点列表, (begin, end, begin, end, ...), 单位s.
   */
  std::vector<float> Detect(const std::vector<char> &wav_data,
                            bool is_end) const;

  ~OnlineVad();

 private:
  class OnlineVadImpl;                     // !< 具体实现类.
  std::unique_ptr<OnlineVadImpl> p_impl_;  // !< 封装了具体实现类的指针
};

/**
 * @brief 离线Vad类.
 * @note 仅支持8k采样频率, 16bit的音频格式.
 */
class OfflineVad {
 public:
  /**
   * @brief 构造.
   * @param vad_yaml vad配置文件路径, 参考vad.yaml.
   * @param inter_threads tensorflow控制op之间并行的线程数, 默认1,
            增大会提升单个文件的速度, 如果文件较多可以使用并行接口.
   * @param intra_threads tensorflow控制op内部并行的线程数, 默认1,
            增大会提升单个文件的速度, 如果文件较多可以使用并行接口.
   */
  explicit OfflineVad(const std::string &vad_yaml, unsigned inter_threads = 1,
                      unsigned intra_threads = 1);

  /**
   * 移动构造函数
   * @param rhs 需要移动的类.
   */
  OfflineVad(OfflineVad &&rhs) noexcept;

  /**
   * @brief 检测Vad端点.
   * @param wav_path wav文件路径.
   * @return 起始终止时间对列表, 单位s, (begin, end).
   */
  std::vector<std::pair<float, float>> Detect(
      const std::string &wav_path) const;

  /**
   * @brief 对多个wav文件并行检测Vad端点.
   * @param wav_paths wav文件路径列表.
   * @return 起始终止时间对列表的列表, 单位s, (begin, end), 与wav文件一一对应.
   */
  std::vector<std::vector<std::pair<float, float>>> DetectMulti(
      const std::vector<std::string> &wav_paths) const;

  ~OfflineVad();

 private:
  class OfflineVadImpl;                     // !< 具体实现类.
  std::unique_ptr<OfflineVadImpl> p_impl_;  // !< 离线Vad对象.
};  // class OfflineVad

/**
 * @brief 在线语种识别类.
 * @note 仅支持8k采样频率, 16bit的音频格式.
 */
class OnlineSlr {
 public:
  /**
   * @brief 构造.
   * @param slr_yaml 语种识别配置文件路径, 参考slr.yaml.
   */
  explicit OnlineSlr(const std::string &slr_yaml);

  /**
   * @brief 重置, 清空所有缓存.
   */
  void Reset();

  /**
   * @brief 检测语种.
   * @param wav_data wav数据, 建议200ms长度, 仅支持8k频率, 16bit的数据块.
   * @param is_end 是否为最后一片.
   * @return 语种信息, cn表示中文, en表示英文, 空字符串表示预测概率小于配置文件中设置的权重.
   */
  std::string Detect(const std::vector<char> &wav_data, bool is_end);

  ~OnlineSlr();

 private:
  class OnlineSlrImpl;                     // !< 具体实现类.
  std::unique_ptr<OnlineSlrImpl> p_impl_;  // !< 封装了具体实现类的指针.
};  // class OnlineSlr

/**
 * @brief Diarization Pipeline类.
 * @note 仅支持8k采样频率, 16bit的音频格式.
 */
class DiarizationPipeline {
 public:
  /**
   * @brief 构造.
   * @param conf_yaml 配置文件路径.
   */
  explicit DiarizationPipeline(const std::string &conf_yaml);

  /**
   * @brief 检测，流程：VAD->SCD->谱聚类.
   * @param wav_path wav文件路径.
   * @return 时间端点(单位s)和说话人ID列表 ((begin1, end1, 1), (begin2, end2, 2), ...).
   */
  std::vector<std::tuple<float, float, unsigned>> Detect(
      const std::string &wav_path) const;

  ~DiarizationPipeline();

 private:
  class DiarizationPipelineImpl;                     // !< 具体实现类.
  std::unique_ptr<DiarizationPipelineImpl> p_impl_;  // !< 离线Vad对象.
};  // class DiarizationPipeline

/**
 * @brief (Depricated) Speaker Identification接口，将在两个版本之后废弃.
 */
class Identification {
 public:
  /**
   * @brief 构造.
   */
  Identification();
  /**
   * @brief 说话人辨别实现.
   * @param speakers_texts 说话人id和文本列表 (1,"你好") (2,"喂") ...
   * @return 说话人身份列表 ("server", "client", ...)
   */
  std::vector<std::string> Detect(
      const std::vector<std::pair<unsigned, std::string>> &speakers_texts);

  ~Identification();

 private:
  class IdentificationImpl;  // !< 具体实现类.
  std::unique_ptr<IdentificationImpl> p_impl_;
};  // class Identification

/**
 * @brief 对输入wav文件和多个模板wav文件进行匹配, 采用cos相似度来计算.
 * @param wav_path 需要匹配的输入wav文件路径.
 * @param tpl_wav_path 模板wav文件路径列表.
 * @return 相似度列表, 对应于输入wav文件和每个模板wav文件的相似度, [0, 1].
 * @note 模板wav文件和需要匹配的wav文件建议使用相同格式的wav文件.
 *       如果输入wav文件时长小于模板wav文件时长, 相似度为0.
 */
std::vector<float> WavMatchCos(const std::string &wav_path,
                               const std::vector<std::string> &tpl_wav_paths);

}  // namespace yispeech
#endif  // YISPEECH_YISPEECH_H_
